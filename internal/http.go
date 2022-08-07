package internal

import (
	"bufio"
	"fmt"
	"net"
)

type http struct {
	conn net.Conn
}

func (h *http) Write(data []byte) error {
	writer := bufio.NewWriter(h.conn)

	if _, err := writer.Write(data); err != nil {
		return err
	}

	return nil
}

func (h *http) Read() ([]byte, error) {
	var buffer []byte

	reader := bufio.NewReader(h.conn)

	if n, err := reader.Read(buffer); err == nil {
		return buffer[:n], nil
	}

	return nil, fmt.Errorf("failed reading")
}
