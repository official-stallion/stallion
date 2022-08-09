package main

import (
	"fmt"
	ponyExpress "github.com/amirhnajafiz/pony-express"
)

func main() {
	client, err := ponyExpress.NewClient("localhost:9090")
	if err != nil {
		panic(err)
	}

	client.Subscribe(func(data []byte) {
		fmt.Println(string(data))
	})

	client.Publish([]byte("Hello world"))

	select {}
}
