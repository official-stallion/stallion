package main

import (
	"log"
	"time"

	"github.com/amirhnajafiz/stallion"
)

func main() {
	client, err := stallion.NewClient("localhost:9090")
	if err != nil {
		panic(err)
	}

	for i := 0; i < 2; i++ {
		go func(j int) {
			cli, er := stallion.NewClient("localhost:9090")
			if er != nil {
				panic(err)
			}

			cli.Subscribe("snapp", func(bytes []byte) {
				log.Printf("%d: %s\n", j, string(bytes))
			})

			cli.Subscribe("snappit", func(bytes []byte) {
				log.Printf("%d second: %s\n", j, string(bytes))
			})

			log.Printf("%d init\n", j)

			select {}
		}(i)
	}

	time.Sleep(1 * time.Second)

	client.Publish("snapp", []byte("Hello"))
	client.Publish("snappit", []byte("Hello to you"))

	time.Sleep(3 * time.Second)

	client.Unsubscribe("snapp")

	select {}
}
