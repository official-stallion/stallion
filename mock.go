package stallion

import (
	"fmt"

	"github.com/official-stallion/stallion/internal"
)

type mockClient struct {
	topics map[string]internal.MessageHandler
}

func NewMockClient() Client {
	return &mockClient{
		topics: make(map[string]internal.MessageHandler),
	}
}

func (m *mockClient) Publish(topic string, data []byte) error {
	if handler, ok := m.topics[topic]; ok {
		handler(data)

		return nil
	}

	return fmt.Errorf("failed to publish")
}

func (m *mockClient) Subscribe(topic string, handler internal.MessageHandler) {
	m.topics[topic] = handler
}

func (m *mockClient) Unsubscribe(topic string) {
	delete(m.topics, topic)
}
