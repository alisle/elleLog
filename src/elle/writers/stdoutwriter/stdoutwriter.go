package StdoutWriter

import (
	"elle/rfc3164"
	"log"
)

var ShowOutput = false
func Process(message *RFC3164.Message) {
	if ShowOutput {
		log.Print(message)
	}
}
