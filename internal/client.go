package internal

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

type client struct {
	connection net.Conn
}

func NewClient(conn net.Conn) *client {
	return &client{
		connection: conn,
	}
}

func (c *client) Publish(data []byte) error {
	writer := bufio.NewWriter(c.connection)
	if _, err := writer.Write(data); err != nil {
		return fmt.Errorf("failed to send: %v\n", err)
	}

	_ = writer.Flush()
	time.Sleep(10 * time.Millisecond)

	fmt.Printf("send %d bytes\n", len(data))

	return nil
}

func (c *client) Subscribe() {
	go func() {
		tmp := make([]byte, 1024)
		for {
			n, err := c.connection.Read(tmp)
			if err != nil {
				if err != io.EOF {
					fmt.Printf("read error: %s\n", err)
				}

				break
			}

			log.Printf("got %d bytes\n", n)
		}
	}()
}
