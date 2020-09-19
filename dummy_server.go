package web_net

type DummyServerHandler struct{}

func (d *DummyServerHandler) Tick(delta float64, events []UserMessage) {
}

func (d *DummyServerHandler) UserConnected(user *User) {

}

func (d *DummyServerHandler) UserDisconnected(user *User) {

}

func (d *DummyServerHandler) AuthenticateTicket(ticketData []byte) (bool, error) {
	return true, nil
}

func (d *DummyServerHandler) AuthenticateConnection(user *User, ticketData []byte) bool {
	return true
}
