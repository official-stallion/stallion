<p align="center">
<img src="assets/logo.png" alt="logo" />
</p>

<p align="center">
<img src="https://img.shields.io/badge/Golang-1.18-66ADD8?style=for-the-badge&logo=go" alt="go version" />
<img src="https://img.shields.io/badge/Version-1.1.1-red?style=for-the-badge&logo=none" alt="version" /><br />
<img src="https://img.shields.io/badge/MacOS-black?style=for-the-badge&logo=apple" alt="version" />
<img src="https://img.shields.io/badge/Linux-white?style=for-the-badge&logo=linux" alt="version" />
<img src="https://img.shields.io/badge/Windows-blue?style=for-the-badge&logo=windows" alt="version" />
</p>

Fast message broker implemented with Golang programming language.<br />
Using no external libraries, just internal Golang libraries.

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
