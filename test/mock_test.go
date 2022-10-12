package test

import (
	"testing"

	"github.com/official-stallion/stallion"
)

func TestMock(t *testing.T) {
	mock := stallion.NewMockClient()
	if mock == nil {
		t.Error("failed to create mock client")
	}

	mock.Subscribe("topic", func(bytes []byte) {
		t.Log("successful subscribe")
	})

	err := mock.Publish("topic", []byte("message"))
	if err != nil {
		t.Error(err)
	}

	mock.Unsubscribe("topic")

	err = mock.Publish("topic", []byte("message"))
	if err == nil {
		t.Error("failed to unsubscribe")
	}
}
