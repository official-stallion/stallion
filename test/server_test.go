package test

import (
	"testing"

	"github.com/official-stallion/stallion"
)

func TestServer(t *testing.T) {
	go func() {
		if err := stallion.NewServer(":6000"); err != nil {
			t.Errorf("server failed to start: %w", err)
		}
	}()

	c, err := stallion.NewClient("localhost:6000")
	if err != nil {
		t.Error(err)
	}

	c.Subscribe("topic", func(bytes []byte) {
		t.Log("success subscribe")
	})

	err = c.Publish("topic", []byte("message"))
	if err != nil {
		t.Error(err)
	}
}