package pony_express

import (
	"fmt"
	"net"

	"github.com/amirhnajafiz/pony-express/internal"
)

type Client interface {
	Publish([]byte) error

	Subscribe()
}

func NewClient(uri string) (Client, error) {
	conn, err := net.Dial("tcp", uri)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to server: %w", err)
	}

	client := internal.NewClient(conn)

	return client, nil
}
