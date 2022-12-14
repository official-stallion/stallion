package test

import (
	"testing"

	"github.com/official-stallion/stallion"
)

// TestServer
// testing stallion server.
func TestServer(t *testing.T) {
	// creating a server on port 6000
	go func() {
		if err := stallion.NewServer(":6000"); err != nil {
			t.Errorf("server failed to start: %v", err)
		}
	}()

	// client does not give a valid url, so we should get error
	c, err := stallion.NewClient("localhost:6000")
	if err == nil {
		t.Error(err)
	}

	// client should connect
	c, err = stallion.NewClient("st://localhost:6000")
	if err != nil {
		t.Error(err)
	}

	// subscribe over a topic
	c.Subscribe("topic", func(bytes []byte) {
		t.Log("success subscribe")
	})

	// we should be able to subscribe
	err = c.Publish("topic", []byte("message"))
	if err != nil {
		t.Error(err)
	}
}

// TestAuthServer
// testing stallion server with auth.
func TestAuthServer(t *testing.T) {
	// creating a stallion server on port 6001 with user and pass
	go func() {
		if err := stallion.NewServer(":6001", "root", "password"); err != nil {
			t.Errorf("server failed to start: %v", err)
		}
	}()

	// client is not authorized we shoud get error
	c, err := stallion.NewClient("st://r:pass@localhost:6001")
	if err == nil {
		t.Error(err)
	}

	// client should connect
	c, err = stallion.NewClient("st://root:password@localhost:6001")
	if err != nil {
		t.Error(err)
	}

	// subscribe over a topic
	c.Subscribe("topic", func(bytes []byte) {
		t.Log("success subscribe")
	})

	// we should be able to subscribe
	err = c.Publish("topic", []byte("message"))
	if err != nil {
		t.Error(err)
	}
}
