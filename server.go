package stallion

import (
	"fmt"
	"net"

	"github.com/official-stallion/stallion/internal"
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

	// handling our clients
	for {
		if conn, er := listener.Accept(); er == nil {
			serve.Handle(conn)
		}
	}
}
