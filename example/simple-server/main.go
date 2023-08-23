package main

import "github.com/official-stallion/stallion"

func main() {
	if err := stallion.NewServer(":9090", 9091); err != nil {
		panic(err)
	}
}
