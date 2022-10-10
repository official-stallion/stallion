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
func urlUnpack(url string) (*url, error) {
	protocolSplit := strings.Split(url, "://")
	if len(protocolSplit) < 2 {
		return nil, fmt.Errorf("invalid uri")
	}

	if protocolSplit[0] != "st" {
		return nil, fmt.Errorf("not using stallion protocol (st://...)")
	}

	return nil, nil
}
