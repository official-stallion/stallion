package internal

import (
	"fmt"
	"net"
)

func New(port string) error {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		return fmt.Errorf("failed to start server: %v", err)
	}

	var (
		receiveChannel chan []byte
		index          = 0
	)

	defer close(receiveChannel)

	b := NewBroker(receiveChannel)
	go b.Start()

	for {
		conn, _ := listener.Accept()

		var temp chan []byte

		w := NewWorker(index, conn, receiveChannel, temp)

		b.Add(temp)

		go w.Start()
	}
}
