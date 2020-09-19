package main

import (
	"fmt"
	"net/http"
	"webnet"
)

type SimpleServer struct {}

func (s *SimpleServer) Tick(delta float64, events []web_net.UserMessage) {
	fmt.Printf("\rTicks: %.4f", delta)

	for _, e := range events {
		fmt.Println(e.Data)
	}
}

func (s *SimpleServer) UserConnected(user *web_net.User) {
	fmt.Println("User Connected", user.ID)
}

func (s *SimpleServer) UserDisconnected(user *web_net.User) {
	fmt.Println("User Disconnected", user.ID)
}

func (s *SimpleServer) AuthenticateTicket(ticketData []byte) (bool, error) {
	return true, nil
}

func (s *SimpleServer) AuthenticateConnection(user *web_net.User, ticketData []byte) bool {
	return true
}

func main() {
	s := web_net.NewServer(&SimpleServer{}, nil)

	err := http.ListenAndServe("0.0.0.0:8899", s)
	if err != nil {
		panic(err)
	}
}
