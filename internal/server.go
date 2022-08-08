package internal

import (
	"log"
	"net"
)

// server is our broker service.
type server struct {
	index         int
	publicChannel chan []byte

	broker *broker
}

// NewServer returns a new broker server.
func NewServer() *server {
	s := &server{
		index: 101,
	}

	s.broker = newBroker(s.publicChannel)
	go s.broker.start()

	return s
}

// Handle will handle the clients.
func (s *server) Handle(conn net.Conn) {
	var temp chan []byte

	w := newWorker(s.index, conn, temp, s.publicChannel)

	log.Printf("new worker %d\n", s.index)

	s.index++
	s.broker.subscribe(temp)

	go w.start()
}
