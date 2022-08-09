package main

func main() {
	if err := stallion.NewServer(":9090"); err != nil {
		panic(err)
	}
}
