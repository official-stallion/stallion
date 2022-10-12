package test

import (
	"testing"

	"github.com/official-stallion/stallion"
)

// TestMock
// testing stallion mock client.
func TestMock(t *testing.T) {
	// creating a mock client
	mock := stallion.NewMockClient()
	if mock == nil {
		t.Error("failed to create mock client")
	}

	// first we subscribe over a topic
	mock.Subscribe("topic", func(bytes []byte) {
		t.Log("successful subscribe")
	})

	// we should be able to publish over that topic
	err := mock.Publish("topic", []byte("message"))
	if err != nil {
		t.Error(err)
	}

	// now we test unsubscribing
	mock.Unsubscribe("topic")

	// we should get an error when we publish over this topic again
	err = mock.Publish("topic", []byte("message"))
	if err == nil {
		t.Error("failed to unsubscribe")
	}
}
