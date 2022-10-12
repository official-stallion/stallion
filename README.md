<p align="center">
<img src="assets/logo.png" alt="logo" />
</p>

<p align="center">
<img src="https://img.shields.io/badge/Golang-1.19-66ADD8?style=for-the-badge&logo=go" alt="go version" />
<img src="https://img.shields.io/badge/Version-1.2.2-red?style=for-the-badge&logo=github" alt="version" /><br />
<img src="https://img.shields.io/badge/MacOS-black?style=for-the-badge&logo=apple" alt="version" />
<img src="https://img.shields.io/badge/Linux-white?style=for-the-badge&logo=linux" alt="version" />
<img src="https://img.shields.io/badge/Windows-blue?style=for-the-badge&logo=windows" alt="version" />
</p>

Fast message broker implemented with Golang programming language.<br />
Using no external libraries, just internal Golang libraries.

## Guide
- [Install Stallion](#how-to-use)
- [Setup Stallion Server](#create-server-in-golang)
- [Using Docker](#create-a-server-with-docker)
- [Stallion Go SDK](#creating-clients)
  - [Subscribe](#subscribe-on-a-topic)
  - [Publish](#publish-over-a-topic)
  - [Unsubscribe](#unsubscribe-from-a-topic)
- [Auth](#creating-a-server-with-auth)
- [Mock](#mock)

## How to use?
Get package:
```shell
go get github.com/official-stallion/stallion@latest
```

Now to set the client up you need to create a **stallion** server.<br />
Stallion server is the message broker server.

### Create server in Golang
```go
package main

import "github.com/official-stallion/stallion"

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

	"github.com/official-stallion/stallion"
)

func main() {
	client, err := stallion.NewClient("st://localhost:9090")
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

## Creating a server with Auth
You can create a Stallion server with username and password for Auth.
```go
package main

import "github.com/official-stallion/stallion"

func main() {
	if err := stallion.NewServer(":9090", "root", "Pa$$word"); err != nil {
		panic(err)
	}
}
```

Now you can connect with username and password set in url.
```go
client, err := stallion.NewClient("st://root:Pa$$word@localhost:9090")
if err != nil {
    panic(err)
}
```

## Mock 
You can create a mock client to create a stallion client sample:
```go
package main

import "github.com/official-stallion/stallion"

func main() {
	client := stallion.NewMockClient()
	
	client.Publish("topic", []byte("message"))
}
```
