package internal

import (
	"log"
	"net"
)

// server is our broker service.
type server struct {
	index  int
	broker *broker
}

// NewServer returns a new broker server.
func NewServer(publicChannel chan []byte) *server {
	s := &server{
		index: 101,
	}

	s.broker = newBroker(publicChannel)
	go s.broker.start()

	return s
}

// Handle will handle the clients.
func (s *server) Handle(conn net.Conn, publicChannel chan []byte) {
	temp := make(chan []byte)

	w := newWorker(s.index, conn, temp, publicChannel)

	log.Printf("new worker %d\n", s.index)

	s.index++
	s.broker.subscribe(temp)

	go w.start()
}
