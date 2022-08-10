package internal

import (
	"bytes"
	"encoding/gob"
)

// Message is what we send between worker and clients.
type Message struct {
	Type  int
	Topic string
	Data  []byte
}

// iom is for send message over tcp
type iom struct {
	Type  int
	Topic string
	Data  string
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
	i := iom{
		Type:  m.Type,
		Topic: m.Topic,
		Data:  string(m.Data),
	}

	var buffer bytes.Buffer

	enc := gob.NewEncoder(&buffer)

	_ = enc.Encode(i)

	return buffer.Bytes()
}

// DecodeMessage will convert array of bytes to Message.
func decodeMessage(buffer []byte) (*Message, error) {
	var m Message
	var i iom

	gob.Register(i)

	dec := gob.NewDecoder(bytes.NewReader(buffer))

	if err := dec.Decode(&i); err != nil {
		if err != bytes.ErrTooLarge {
			return nil, err
		}
	}

	m.Type = i.Type
	m.Topic = i.Topic
	m.Data = []byte(i.Data)

	return &m, nil
}
