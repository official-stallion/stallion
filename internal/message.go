package internal

import (
	"encoding/json"
)

// Message is what we send between worker and clients.
type message struct {
	Type  int    `json:"type"`
	Topic string `json:"topic"`
	Data  []byte `json:"data"`
}

// NewMessage generates a new message type.
func newMessage(t int, topic string, data []byte) message {
	return message{
		Type:  t,
		Topic: topic,
		Data:  data,
	}
}

// EncodeMessage will convert message to array of bytes.
func encodeMessage(m message) []byte {
	bytes, _ := json.Marshal(m)

	return bytes
}

// DecodeMessage will convert array of bytes to Message.
func decodeMessage(buffer []byte) (*message, error) {
	var m message

	if err := json.Unmarshal(buffer, &m); err != nil {
		return nil, err
	}

	return &m, nil
}
