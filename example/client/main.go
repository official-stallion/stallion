package main

import (
	"time"

	ponyExpress "github.com/amirhnajafiz/pony-express"
)

func main() {
	client, err := ponyExpress.NewClient("localhost:9090")
	if err != nil {
		panic(err)
	}

	client.Subscribe()

	_ = client.Publish([]byte("Hello world 1"))
	_ = client.Publish([]byte("Hello world 2"))
	_ = client.Publish([]byte("Hello world 3"))

	time.Sleep(5 * time.Second)
}
