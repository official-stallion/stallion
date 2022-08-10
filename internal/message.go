package internal

import (
	"encoding/json"
)

// Message is what we send between worker and clients.
type Message struct {
	Type  int    `json:"type"`
	Topic string `json:"topic"`
	Data  []byte `json:"data"`
}

// NewMessage generates a new message type.
func newMessage(t int, data []byte) Message {
	return Message{
		Type: t,
		Data: data,
	}
}

// EncodeMessage will convert message to array of bytes.
func encodeMessage(m Message) []byte {
	bytes, _ := json.Marshal(m)

	return bytes
}

// DecodeMessage will convert array of bytes to Message.
func decodeMessage(bytes []byte) (*Message, error) {
	var m Message

	if err := json.Unmarshal(bytes, &m); err != nil {
		return nil, err
	}

	return &m, nil
}
