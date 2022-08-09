package main

import (
	"fmt"
	"time"
)

func main() {
	client, err := stallion.NewClient("localhost:9090")
	if err != nil {
		panic(err)
	}

	client.Subscribe(func(data []byte) {
		fmt.Println(string(data))
	})

	client.Publish([]byte("Hello"))

	time.Sleep(3 * time.Second)

	client.Unsubscribe()

	select {}
}
