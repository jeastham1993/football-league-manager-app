package infrastructure

import (
	"log"
)

// Logger encapsulates code for logging data.
type Logger struct{}

// Log writes message to the console.
func (logger Logger) Log(args ...interface{}) {
	log.Println(args...)
}
