package internal

import "encoding/json"

// Message is what we send between worker and clients.
type Message struct {
	Type int    `json:"type"`
	Data []byte `json:"message"`
}

// NewMessage generates a new message type.
func NewMessage(t int, data []byte) Message {
	return Message{
		Type: t,
		Data: data,
	}
}

// EncodeMessage will convert message to array of bytes.
func EncodeMessage(m Message) []byte {
	bytes, _ := json.Marshal(m)

	return bytes
}

// DecodeMessage will convert array of bytes to Message.
func DecodeMessage(bytes []byte) (*Message, error) {
	var m Message

	if err := json.Unmarshal(bytes, &m); err != nil {
		return nil, err
	}

	return &m, nil
}
