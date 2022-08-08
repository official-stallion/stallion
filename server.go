package pony_express

import (
	"fmt"
	"log"
	"net"

	"github.com/amirhnajafiz/pony-express/internal"
)

func NewServer(port string) error {
	channel := make(chan []byte)

	serve := internal.NewServer(channel)

	listener, err := net.Listen("tcp", port)
	if err != nil {
		return fmt.Errorf("failed to start server: %v", err)
	}

	log.Printf("start broker server on %s ...\n", port)

	for {
		conn, _ := listener.Accept()
		serve.Handle(conn, channel)
	}
}
