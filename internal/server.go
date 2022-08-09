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
func NewServer(publicChannel chan []byte, status chan int) *server {
	s := &server{
		index: 101,
	}

	s.broker = newBroker(publicChannel, status)
	go s.broker.start()

	return s
}

// Handle will handle the clients.
func (s *server) Handle(conn net.Conn, publicChannel chan []byte, status chan int) {
	w := newWorker(s.index, conn, make(chan []byte), publicChannel, status)

	log.Printf("new worker %d\n", s.index)

	s.index++
	s.broker.subscribe(w.sendChannel, w.id)

	go w.start()
}
