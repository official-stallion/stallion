package internal

import (
	"fmt"
	"io"
	"net"
)

// network handles the tcp requests.
type network struct {
	connection net.Conn
}

// send data over tcp.
func (n *network) send(data []byte) error {
	if _, err := n.connection.Write(data); err != nil {
		return fmt.Errorf("failed to send: %v\n", err)
	}

	return nil
}

// get data from tcp.
func (n *network) get(buffer []byte) ([]byte, error) {
	bytes, err := n.connection.Read(buffer)
	if err != nil {
		if err != io.EOF {
			logError("network read error", err)
		}

		return nil, err
	}

	return buffer[:bytes], nil
}
