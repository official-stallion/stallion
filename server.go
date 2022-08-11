package stallion

import (
	"fmt"
	"log"
	"net"

	"github.com/amirhnajafiz/stallion/internal"
)

type Server interface {
	// Handle method generates a new worker for clients.
	Handle(conn net.Conn)
}

// NewServer creates a new broker server on given port.
func NewServer(port string) error {
	// creating a new server
	serve := internal.NewServer()

	// listen over a port
	listener, err := net.Listen("tcp", port)
	if err != nil {
		return fmt.Errorf("failed to start server: %v", err)
	}

	log.Printf("start broker server on %s ...\n", port)

	// handling our clients
	for {
		if conn, er := listener.Accept(); er == nil {
			serve.Handle(conn)
		} else {
			log.Printf("error in client accept: %v\n", er)
		}
	}
}
