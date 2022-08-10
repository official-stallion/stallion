<h1 align="center">
Stallion
</h1>

Message broker with Golang.

## How to use?
Get package:
```shell
go get github.com/amirhnajafiz/stallion
```

### Client
```go
package main

import (
	"fmt"
	"time"

	"github.com/amirhnajafiz/stallion"
)

func main() {
	client, err := stallion.NewClient("localhost:9090")
	if err != nil {
		panic(err)
	}

	client.Subscribe("topic", func(data []byte) {
		fmt.Println(string(data))
	})

	client.Publish("topic", []byte("Hello"))

	time.Sleep(3 * time.Second)
}
```

### Server
```go
package main

import "github.com/amirhnajafiz/stallion"

func main() {
	if err := stallion.NewServer(":9090"); err != nil {
		panic(err)
	}
}
```
