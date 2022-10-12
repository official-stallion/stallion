package stallion

import "github.com/official-stallion/stallion/internal"

type mockClient struct{}

func NewMockClient() Client {
	return &mockClient{}
}

func (m *mockClient) Publish(topic string, data []byte) error {

}

func (m *mockClient) Subscribe(topic string, handler internal.MessageHandler) {

}

func (m *mockClient) Unsubscribe(topic string) {

}
