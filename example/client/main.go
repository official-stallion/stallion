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
	client.Publish([]byte("Hello world"))

	time.Sleep(3 * time.Second)
}
