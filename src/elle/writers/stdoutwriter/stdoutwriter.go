package StdoutWriter

import (
	"elle/rfc3164"
	"log"
)

func Process(message *RFC3164.Message) {
	log.Print(message)
}
