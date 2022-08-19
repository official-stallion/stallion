package internal

import (
	"log"
	"time"
)

// logInfo creates an info log.
func logInfo(msg string) {
	t := time.Now().Format(time.RFC3339)

	log.Printf("[info][%s] %s\n", t, msg)
}

// logError creates an error log.
func logError(msg string, e error) {
	t := time.Now().Format(time.RFC3339)

	log.Printf("[error][%s] %s: %v\n", t, msg, e)
}
