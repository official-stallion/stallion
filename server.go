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
func NewServer(port string, auth ...string) error {
	// get authentication options
	var (
		user string
		pass string
	)

	// setting the authentication options
	if len(auth) > 1 {
		user = auth[0]
		pass = auth[1]
	}

	// creating a new server
	serve := internal.NewServer(user, pass)

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
