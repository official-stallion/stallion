package main

import (
	"log"

	"github.com/amirhnajafiz/stallion"
)

func main() {
	for i := 0; i < 2; i++ {
		go func(j int) {
			cli, er := stallion.NewClient("localhost:9090")
			if er != nil {
				panic(er)
			}

			cli.Subscribe("topic1", func(bytes []byte) {
				log.Printf("%d: %s\n", j, string(bytes))
			})

			cli.Subscribe("topic2", func(bytes []byte) {
				log.Printf("%d second: %s\n", j, string(bytes))
			})

			select {}
		}(i)
	}

	client, err := stallion.NewClient("localhost:9090")
	if err != nil {
		panic(err)
	}

	client.Publish("topic1", []byte("Hello from 1"))
	client.Publish("topic2", []byte("Hello from 2"))

	select {}
}
