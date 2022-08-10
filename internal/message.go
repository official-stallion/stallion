package internal

import (
	"encoding/json"
)

// Message is what we send between worker and clients.
type Message struct {
	Type  int
	Topic string
	Data  []byte
}

// jsonMessage is for send message over tcp
type jsonMessage struct {
	Type  int    `json:"type"`
	Topic string `json:"topic"`
	Data  string `json:"data"`
}

// NewMessage generates a new message type.
func newMessage(t int, topic string, data []byte) Message {
	return Message{
		Type:  t,
		Topic: topic,
		Data:  data,
	}
}

// EncodeMessage will convert message to array of bytes.
func encodeMessage(m Message) []byte {
	j := jsonMessage{
		Type:  m.Type,
		Topic: m.Topic,
		Data:  string(m.Data),
	}

	bytes, _ := json.Marshal(j)

	return bytes
}

// DecodeMessage will convert array of bytes to Message.
func decodeMessage(buffer []byte) (*Message, error) {
	var (
		m Message
		i jsonMessage
	)

	if err := json.Unmarshal(buffer, &i); err != nil {
		return nil, err
	}

	m.Type = i.Type
	m.Topic = i.Topic
	m.Data = []byte(i.Data)

	return &m, nil
}
