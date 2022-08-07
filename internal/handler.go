package internal

import (
	"bufio"
	"fmt"
	"net"
)

// handler handles a client for communication.
type handler struct {
	conn net.Conn
}

// Write will write a data into client socket.
func (h *handler) Write(data []byte) error {
	writer := bufio.NewWriter(h.conn)

	if _, err := writer.Write(data); err != nil {
		return err
	}

	return nil
}

// Read will read a data from client socket.
func (h *handler) Read() ([]byte, error) {
	var buffer []byte

	reader := bufio.NewReader(h.conn)

	if n, err := reader.Read(buffer); err == nil {
		return buffer[:n], nil
	}

	return nil, fmt.Errorf("failed reading")
}
