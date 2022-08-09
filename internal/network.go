package internal

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

// network handles the tcp requests.
type network struct {
	connection net.Conn
}

// send data over tcp.
func (n *network) send(data []byte) error {
	writer := bufio.NewWriter(n.connection)
	if _, err := writer.Write(data); err != nil {
		return fmt.Errorf("failed to send: %v\n", err)
	}

	_ = writer.Flush()

	return nil
}

// get data from tcp.
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