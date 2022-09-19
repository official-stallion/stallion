package main

import (
	"fmt"
	"time"

	"github.com/official-stallion/stallion"
)

func main() {
	client, err := stallion.NewClient("localhost:9090")
	if err != nil {
		panic(err)
	}

	client.Subscribe("topic", func(data []byte) {
		fmt.Println(string(data))
	})

	client.Publish("topic", []byte("Hello"))

	time.Sleep(3 * time.Second)
}
