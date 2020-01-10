package infrastructure

import (
	"log"
)

// Logger encapsulates code for logging data.
type Logger struct{}

// Log writes message to the console.
func (logger Logger) Log(message string) {
	log.Println(message)
}
