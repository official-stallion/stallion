package main

import (
	"os"

	"github.com/amirhnajafiz/stallion"
)

const (
	defaultPort = "7025"
)

func main() {
	port := defaultPort
	if value, ok := os.LookupEnv("SERVER_PORT"); ok {
		port = value
	}

	if err := stallion.NewServer(":" + port); err != nil {
		panic(err)
	}
}
