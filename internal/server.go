package internal

import (
	"net"
)

// server is our broker service.
type server struct {
	prefix int
	broker *broker
}

// NewServer returns a new broker server.
func NewServer(public chan []byte, status chan WorkChan) *server {
	s := &server{
		prefix: 101,
	}

	// setting up the server broker and starting it
	s.broker = newBroker(public, status)
	go s.broker.start()

	return s
}

// Handle will handle the clients.
func (s *server) Handle(conn net.Conn, public chan []byte, status chan WorkChan) {
	w := newWorker(s.prefix, conn, make(chan []byte), public, status)

	s.prefix++
	s.broker.subscribe(w.sendChannel, w.id)

	go w.start()
}
