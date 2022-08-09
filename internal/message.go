package internal

import (
	"encoding/json"
)

// Message is what we send between worker and clients.
type message struct {
	Type int    `json:"type"`
	Data []byte `json:"data"`
}

// NewMessage generates a new message type.
func newMessage(t int, data []byte) message {
	return message{
		Type: t,
		Data: data,
	}
}

// EncodeMessage will convert message to array of bytes.
func encodeMessage(m message) []byte {
	bytes, _ := json.Marshal(m)

	return bytes
}

// DecodeMessage will convert array of bytes to Message.
func decodeMessage(bytes []byte) (*message, error) {
	var m message

	if err := json.Unmarshal(bytes, &m); err != nil {
		return nil, err
	}

	return &m, nil
}
