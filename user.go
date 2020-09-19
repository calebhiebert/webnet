package web_net

import "sync"

type User struct {
	ID         string      `json:"id" msgpack:"id"`
	Data       interface{} `json:"data" msgpack:"data"`      // Data stored locally on the server
	ClientData interface{} `json:"client_data" msgpack:"cd"` // Data synchronized with the client
	TicketData []byte      `json:"ticket_data" msgpack:"td"` // Connection ticket data made available on the user

	mx       sync.Mutex  // Any operations involving the user data should use this mutex
	writer   chan []byte // Any messages to be sent to the user's client should be put in this channel
	messages []Message   // All messages received from this user since the last tick
}
