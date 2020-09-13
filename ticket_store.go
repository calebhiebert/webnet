package web_net

import "errors"

var (
	ErrTicketNotFound = errors.New("ticket not found")
)

// ITicketStore is a helper to store websocket tickets
type ITicketStore interface {

	// StoreTicket should take a ticket id, and some data and persist it somewhere for later retrieval
	StoreTicket(ticket string, data []byte) error

	// RetrieveTicket should take a ticket id and return the previously stored ticket data
	// (or ErrTicketNotFound if the ticket did not exist)
	RetrieveTicket(ticket string) ([]byte, error)
}

// MemoryTicketStore will store ticket data in memory, not safe for production
type MemoryTicketStore struct {
	tickets map[string][]byte
}

func (m *MemoryTicketStore) StoreTicket(ticket string, data []byte) error {
	if m.tickets == nil {
		m.tickets = map[string][]byte{}
	}

	m.tickets[ticket] = data

	return nil
}

func (m *MemoryTicketStore) RetrieveTicket(ticket string) ([]byte, error) {
	if m.tickets == nil {
		return nil, ErrTicketNotFound
	}

	if data, ok := m.tickets[ticket]; ok {
		return data, nil
	}

	return nil, ErrTicketNotFound
}
