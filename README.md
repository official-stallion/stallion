<p align="center">
<img src="assets/logo.png" alt="logo" />
</p>

<p align="center">
<img src="https://img.shields.io/badge/Golang-1.18-66ADD8?style=for-the-badge&logo=go" alt="go version" />
<img src="https://img.shields.io/badge/Version-1.1.2-red?style=for-the-badge&logo=none" alt="version" /><br />
<img src="https://img.shields.io/badge/MacOS-black?style=for-the-badge&logo=apple" alt="version" />
<img src="https://img.shields.io/badge/Linux-white?style=for-the-badge&logo=linux" alt="version" />
<img src="https://img.shields.io/badge/Windows-blue?style=for-the-badge&logo=windows" alt="version" />
</p>

Fast message broker implemented with Golang programming language.<br />
Using no external libraries, just internal Golang libraries.

## How to use?
Get package:
```shell
go get github.com/amirhnajafiz/stallion@latest
```

Now to set the client up you need to create a **stallion** server.<br />
Stallion server is the message broker server.

### Create server in Golang
```go
package main

import "github.com/amirhnajafiz/stallion"

func main() {
	if err := stallion.NewServer(":9090"); err != nil {
		panic(err)
	}
}
```

### Create a server with docker
Check the docker [documentation](./docker/README.md) for stallion server.

### Creating Clients
You can connect to stallion server like the example below:
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
}
```

### Subscribe on a topic
```go
client.Subscribe("topic", func(data []byte) {
    // any handler that you want
    fmt.Println(string(data))
})
```

### Publish over a topic
```go
client.Publish("topic", []byte("Hello"))
```

### Unsubscribe from a topic
```go
client.Unsubscribe("topic")
```

