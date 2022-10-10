package main

import (
	"log"
	"time"

	"github.com/official-stallion/stallion"
)

func main() {
	client, err := stallion.NewClient("st://localhost:9090")
	if err != nil {
		panic(err)
	}

	for i := 0; i < 2; i++ {
		go func(j int) {
			cli, er := stallion.NewClient("st://localhost:9090")
			if er != nil {
				log.Fatal(er)
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

	time.Sleep(2 * time.Second)

	client.Publish("topic1", []byte("Hello from 1"))
	client.Publish("topic2", []byte("Hello from 2"))

	select {}
}
