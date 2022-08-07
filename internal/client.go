package internal

import (
	"fmt"
	"log"
	"net"
)

type client struct {
	handler handler
}

func NewClient(conn net.Conn) *client {
	return &client{
		handler: handler{
			conn: conn,
		},
	}
}

func (c *client) Send(data []byte) {
	if err := c.handler.write(data); err != nil {
		log.Fatalf("failed to send: %v\n", err)
	}
}

func (c *client) Receive() {
	for {
		data, err := c.handler.read()
		if err == nil {
			fmt.Println(data)
		}
	}
}
