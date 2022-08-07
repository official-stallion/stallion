package internal

import (
	"fmt"
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

func (c *client) Publish(data []byte) error {
	if err := c.handler.write(data); err != nil {
		return fmt.Errorf("failed to send: %v\n", err)
	}

	return nil
}

func (c *client) Subscribe() {
	go func() {
		for {
			data, err := c.handler.read()
			if err == nil {
				fmt.Println(data)
			}
		}
	}()
}
