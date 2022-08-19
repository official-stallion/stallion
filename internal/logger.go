package internal

import (
	"log"
	"strings"
)

// logInfo creates an info log.
func logInfo(msg string, fields ...string) {
	if len(fields) > 0 {
		log.Printf("[info] %s: %s\n", msg, strings.Join(fields, ", "))
	} else {
		log.Printf("[info] %s\n", msg)
	}
}

// logError creates an error log.
func logError(msg string, e error, fields ...string) {
	if len(fields) > 0 {
		log.Printf("[error] %s: %s: %v\n", msg, strings.Join(fields, ", "), e)
	} else {
		log.Printf("[error] %s: %v\n", msg, e)
	}
}
