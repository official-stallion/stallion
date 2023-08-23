package main

import "github.com/official-stallion/stallion"

func main() {
	if err := stallion.NewServer(":9090", 9091, "root", "Pa$$word"); err != nil {
		panic(err)
	}
}
