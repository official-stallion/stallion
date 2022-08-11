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
func NewServer() *server {
	s := &server{
		prefix: 101,
	}

	// setting up the server broker and starting it
	s.broker = newBroker(
		make(chan Message),
		make(chan SubscribeChannel),
		make(chan UnsubscribeChannel),
		make(chan int),
	)
	go s.broker.start()

	return s
}

// Handle will handle the clients.
func (s *server) Handle(conn net.Conn) {
	w := newWorker(
		s.prefix,
		conn,
		make(chan Message),
		s.broker.receiveChannel,
		s.broker.subscribeChannel,
		s.broker.unsubscribeChannel,
		s.broker.terminateChannel,
	)

	s.prefix++

	go w.start()
}
