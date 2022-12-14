package internal

import (
	"fmt"
	"net"
)

// server is our broker service.
type server struct {
	auth auth

	prefix int
	broker *broker
}

// NewServer returns a new broker server.
func NewServer(user string, pass string) *server {
	s := &server{
		auth: auth{
			username: user,
			password: pass,
		},
		prefix: 101,
	}

	// setting up the server broker and starting it
	s.broker = newBroker(
		make(chan message),
		make(chan subscribeChannel),
		make(chan unsubscribeChannel),
		make(chan int),
	)
	go s.broker.start()

	return s
}

// Handle will handle the clients.
func (s *server) Handle(conn net.Conn) {
	w := newWorker(
		s.prefix,
		s.auth,
		conn,
		make(chan message),
		s.broker.receiveChannel,
		s.broker.subscribeChannel,
		s.broker.unsubscribeChannel,
		s.broker.terminateChannel,
	)

	logInfo("new client joined", fmt.Sprintf("id=%d", s.prefix))

	s.prefix++

	go w.start()
}
