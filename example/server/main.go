package main

import "github.com/amirhnajafiz/stallion"

func main() {
	if err := stallion.NewServer(":9090"); err != nil {
		panic(err)
	}
}
