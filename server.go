package pony_express

import (
	"fmt"
	"log"
	"net"

	"github.com/amirhnajafiz/pony-express/internal"
)

func NewServer(port string) error {
	serve := internal.NewServer()

	listener, err := net.Listen("tcp", port)
	if err != nil {
		return fmt.Errorf("failed to start server: %v", err)
	}

	log.Printf("start broker server on %s ...\n", port)

	for {
		conn, _ := listener.Accept()
		serve.Handle(conn)
	}
}
