package internal

import (
	"fmt"
	"net"
)

type server struct{}

func NewServer() *server {
	return &server{}
}

func (s *server) start(port string) error {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		return fmt.Errorf("failed to start server: %v", err)
	}

	var (
		publicChannel chan []byte
		index         = 0
	)

	defer close(publicChannel)

	broker := newBroker(publicChannel)
	go broker.start()

	for {
		conn, _ := listener.Accept()

		var temp chan []byte

		w := newWorker(index, conn, publicChannel, temp)

		broker.subscribe(temp)

		go w.start()
	}
}
