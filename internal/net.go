package internal

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

type network struct {
	connection net.Conn
}

func (n *network) send(data []byte) error {
	writer := bufio.NewWriter(n.connection)
	if _, err := writer.Write(data); err != nil {
		return fmt.Errorf("failed to send: %v\n", err)
	}

	_ = writer.Flush()

	return nil
}

func (n *network) get(buffer []byte) ([]byte, error) {
	bytes, err := n.connection.Read(buffer)
	if err != nil {
		if err != io.EOF {
			log.Printf("read error: %s\n", err)
		}

		return nil, err
	}

	return buffer[:bytes], nil
}
