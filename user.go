package web_net

type User struct {
	ID         string      `json:"id" msgpack:"id"`
	Data       interface{} `json:"data" msgpack:"data"`      // Data stored locally on the server
	ClientData interface{} `json:"client_data" msgpack:"cd"` // Data synchronized with the client

	writer chan []byte
}
