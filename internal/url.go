package internal

import (
	"fmt"
	"strings"
)

// url
// each url contains of the following parts:
// - host
// - port
type url struct {
	host string
	port string
}

// urlUnpack
// manages to create url struct from url string.
func urlUnpack(inputUrl string) (*url, error) {
	protocolSplit := strings.Split(inputUrl, "://")
	if len(protocolSplit) < 2 {
		return nil, fmt.Errorf("invalid uri")
	}

	if protocolSplit[0] != "st" {
		return nil, fmt.Errorf("not using stallion protocol (st://...)")
	}

	hostInformation := strings.Split(protocolSplit[1], ":")
	if len(hostInformation) < 2 {
		return nil, fmt.Errorf("server ip or port is not given")
	}

	return &url{
		host: hostInformation[0],
		port: hostInformation[1],
	}, nil
}
