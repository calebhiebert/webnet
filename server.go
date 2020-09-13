package web_net

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	gonanoid "github.com/matoous/go-nanoid"
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
}

type IServerHandler interface {
	AuthenticateTicket(ticketData []byte) (bool, error)
	AuthenticateConnection(user User, ticketData []byte) bool
}

func NewServer(handler IServerHandler) *Server {
	return &Server{
		handler:     handler,
		ticketStore: &MemoryTicketStore{},
	}
}

func (s *Server) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/ticket" && req.Method == http.MethodPost {
		s.handleTicket(res, req)
	} else if req.URL.Path == "/ws" {
		s.handleWS(res, req)
	}
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
	_, err := s.ticketStore.RetrieveTicket(ticket)
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

	// TODO do something with ticket data

	conn, err := wsupgrader.Upgrade(res, req, nil)
	if err != nil {
		fmt.Println("Failed to upgrade websocket", err)
		return
	}

	writer := make(chan []byte, 5)
	conMX := sync.Mutex{}

	// Message writer goroutine
	go func() {
		for m := range writer {
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
			t, _, err := conn.ReadMessage()
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
		}
	}()
}
