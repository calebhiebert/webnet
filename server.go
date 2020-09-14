package web_net

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	gonanoid "github.com/matoous/go-nanoid"
	"github.com/vmihailenco/msgpack/v4"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Server struct {
	ticketStore ITicketStore
	handler     IServerHandler

	userMX         sync.Mutex
	connectedUsers map[string]*User
	simpleCORS     bool
}

type Message struct {
	Type string      `json:"t" msgpack:"t"`
	Data interface{} `json:"d" msgpack:"d"`
}

type IServerHandler interface {
	AuthenticateTicket(ticketData []byte) (bool, error)
	AuthenticateConnection(user *User, ticketData []byte) bool
	UserConnected(user *User)
	UserDisconnected(user *User)
}

func NewServer(handler IServerHandler) *Server {
	return &Server{
		handler:        handler,
		simpleCORS:     true,
		ticketStore:    &MemoryTicketStore{},
		connectedUsers: map[string]*User{},
	}
}

func (s *Server) SetSimpleCORS(b bool) {
	s.simpleCORS = b
}

func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/ticket" {
		if s.simpleCORS {
			s.writeCORSHeaders(&w)
		}

		if req.Method == "POST" {
			s.handleTicket(w, req)
		} else if req.Method == "OPTIONS" && s.simpleCORS {
			return
		}
	} else if req.URL.Path == "/ws" {
		if s.simpleCORS {
			s.writeCORSHeaders(&w)
		}

		s.handleWS(w, req)
	}
}

func (s *Server) writeCORSHeaders(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding")
}

func (s *Server) handleTicket(res http.ResponseWriter, req *http.Request) {
	respond := func(code int, data map[string]string) {
		jsb, _ := json.Marshal(data)
		res.WriteHeader(code)
		_, _ = res.Write(jsb)
	}

	ticketData, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("Failed to read body", err)
		respond(http.StatusOK, map[string]string{
			"error": err.Error(),
		})
		return
	}

	authorized, err := s.handler.AuthenticateTicket(ticketData)
	if err != nil {
		respond(http.StatusUnauthorized, map[string]string{"status": "not authorized", "error": err.Error()})
		return
	}

	if !authorized {
		respond(http.StatusUnauthorized, map[string]string{"error": "not authorized"})
		return
	}

	// Generate ticket id and store ticket data
	ticketID, err := gonanoid.Nanoid(32)
	if err != nil {
		fmt.Println("failed to generate ticket id", err)
		respond(http.StatusInternalServerError, map[string]string{"error": "something went wrong"})
		return
	}

	err = s.ticketStore.StoreTicket(ticketID, ticketData)
	if err != nil {
		fmt.Println("failed to store ticket", err)
		respond(http.StatusInternalServerError, map[string]string{"error": "something went wrong"})
	}

	respond(http.StatusOK, map[string]string{"ticket": ticketID})
}

func (s *Server) handleWS(res http.ResponseWriter, req *http.Request) {
	// Get the ticket from the query parameter
	ticket := req.URL.Query().Get("ticket")

	// Attempt to get the ticket data from the store
	ticketData, err := s.ticketStore.RetrieveTicket(ticket)
	if err != nil {
		if err == ErrTicketNotFound {
			res.WriteHeader(http.StatusUnauthorized)
			_, _ = res.Write([]byte("invalid ticket"))
			return
		} else {
			res.WriteHeader(http.StatusInternalServerError)
			_, _ = res.Write([]byte("something went wrong"))
			return
		}
	}

	// Create user object and add it to server
	uid, _ := gonanoid.Nanoid()

	user := &User{
		ID:         uid,
		TicketData: ticketData,
		Data:       nil,
		ClientData: nil,
		writer:     make(chan []byte, 5),
		reader:     make(chan []byte, 5),
	}

	ticketDataValid := s.handler.AuthenticateConnection(user, ticketData)

	if !ticketDataValid {
		jsb, _ := json.Marshal(map[string]interface{}{"error": "not authorized"})

		res.WriteHeader(http.StatusForbidden)
		_, _ = res.Write(jsb)
		return
	}

	conn, err := wsupgrader.Upgrade(res, req, nil)
	if err != nil {
		fmt.Println("Failed to upgrade websocket", err)
		return
	}

	conMX := sync.Mutex{}

	// Message writer goroutine
	go func() {
		for m := range user.writer {
			conMX.Lock()
			err := conn.WriteMessage(websocket.BinaryMessage, m)
			if err != nil {
				fmt.Println("Failed to write message", err)
			}
			conMX.Unlock()
		}
	}()

	// Keepalive goroutine
	go func() {
		for {
			time.Sleep(5 * time.Second)

			conMX.Lock()
			err = conn.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(5*time.Second))
			if err != nil {
				if err == websocket.ErrCloseSent {
					conMX.Unlock()
					break
				} else if strings.Contains(err.Error(), "use of closed network connection") {
					conMX.Unlock()
					break
				} else {
					fmt.Println("Error writing keepalive", err)
				}
			}
			conMX.Unlock()
		}
	}()

	// Message reader goroutine
	go func() {
		for {
			t, m, err := conn.ReadMessage()
			if err != nil {
				switch err.(type) {
				case *websocket.CloseError:
					closeErr := err.(*websocket.CloseError)

					switch closeErr.Code {
					case websocket.CloseNormalClosure, websocket.CloseGoingAway, websocket.CloseAbnormalClosure, websocket.CloseNoStatusReceived:
						// Don't do anything
					default:
						fmt.Println("Error reading websocket message", err)
					}

					_ = conn.Close()
				case *net.OpError:
					fmt.Println("websocket net error ", err)
				default:
					fmt.Println("websocket unknown error ", err)
				}

				break
			}

			// Check to make sure the message is a binary message
			if t != websocket.BinaryMessage {
				fmt.Println("Incorrect message type", t)
			}

			user.reader <- m
		}

		// Once this point is reached, the connection has been closed
		s.userDisconnected(user)
	}()

	s.addUser(user)
}

func (s *Server) addUser(user *User) {
	s.userMX.Lock()
	s.connectedUsers[user.ID] = user
	s.userMX.Unlock()

	s.handler.UserConnected(user)

	user.mx.Lock()
	user.ClientData = map[string]interface{}{
		"_id": user.ID,
	}
	user.mx.Unlock()

	s.syncUserState(user)
	s.syncGameStateToUser(user)
}

func (s *Server) syncUserState(user *User) {
	user.mx.Lock()
	userData := user.ClientData
	user.mx.Unlock()

	// Send the user's own state to them
	udBytes, err := msgpack.Marshal(Message{Type: "ud_state_sync", Data: userData})
	if err != nil {
		fmt.Println("Failed to marshal user state")
	} else {
		user.writer <- udBytes
	}
}

func (s *Server) syncGameStateToUser(user *User) {
	// Send the current game state to the user
	gsBytes, err := msgpack.Marshal(Message{Type: "gs_sync", Data: "imaginary state"})
	if err != nil {
		fmt.Println("Failed to marshal game state")
	} else {
		user.writer <- gsBytes
	}
}

func (s *Server) userDisconnected(user *User) {
	s.userMX.Lock()
	delete(s.connectedUsers, user.ID)
	s.userMX.Unlock()

	s.handler.UserDisconnected(user)
}
