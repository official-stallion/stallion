package stallion

import (
	"fmt"
	"github.com/official-stallion/stallion/internal"
)

// mockClient
// mocks stallion client.
type mockClient struct {
	// list of the topics with their handlers
	topics map[string]internal.MessageHandler
}

// NewMockClient
// creates a new mock client.
func NewMockClient() Client {
	return &mockClient{
		topics: make(map[string]internal.MessageHandler),
	}
}

// Publish
// send messages over mock client.
func (m *mockClient) Publish(topic string, data []byte) error {
	if handler, ok := m.topics[topic]; ok {
		handler(data)

		return nil
	}

	return fmt.Errorf("failed to publish")
}

// Subscribe
// over a topic.
func (m *mockClient) Subscribe(topic string, handler internal.MessageHandler) {
	m.topics[topic] = handler
}

// Unsubscribe
// from a topic.
func (m *mockClient) Unsubscribe(topic string) {
	delete(m.topics, topic)
}
