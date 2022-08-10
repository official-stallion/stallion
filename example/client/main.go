package main

import (
	"fmt"
	"time"

	"github.com/amirhnajafiz/stallion"
)

func main() {
	client, err := stallion.NewClient("localhost:9090")
	if err != nil {
		panic(err)
	}

	client.Subscribe("snapp", func(data []byte) {
		fmt.Println(string(data))
	})

	client.Publish("snapp", []byte("Hello"))

	time.Sleep(3 * time.Second)

	client.Unsubscribe("snapp")

	select {}
}
