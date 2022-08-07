package main

import ponyExpress "github.com/amirhnajafiz/pony-express"

func main() {
	if err := ponyExpress.NewServer(":9090"); err != nil {
		panic(err)
	}
}
