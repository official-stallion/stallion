<p align="center">
<img src="assets/logo.png" alt="logo" />
</p>

<p align="center">
<img src="https://img.shields.io/badge/Golang-1.19-66ADD8?style=for-the-badge&logo=go" alt="go version" />
<img src="https://img.shields.io/badge/Version-1.3.0-red?style=for-the-badge&logo=github" alt="version" /><br />
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
- [Auth](#creating-a-server-with-auth)

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
