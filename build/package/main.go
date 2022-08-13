package main

import (
	"log"
	"os"
)

const (
	defaultPort = "7025"
)

func main() {
	port := defaultPort
	if value, ok := os.LookupEnv("SERVER_PORT"); ok {
		port = value
	}

	log.Println(":" + port)
}
