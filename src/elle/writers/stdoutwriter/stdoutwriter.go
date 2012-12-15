package StdoutWriter

import (
	"elle/rfc3164"
    "elle/processors"
	"log"
    "sort"
)

var ShowOutput = false
func ProcessMessage(message *RFC3164.Message) {
	if ShowOutput {
		log.Print(message)
	}
}

func Process(event Processors.Event) {
    if ShowOutput {
        DumpEvent(event)
    }
}
func DumpEvent(event Processors.Event) {
    keys := make([]string, len(event))
    i := 0
    for key, _ := range event {
        keys[i] = key
        i++
    }
    sort.Strings(keys)

    log.Print("New Event:")
    for _, key := range keys {
        log.Print("\t", key, " : ", event[key])
    }

    log.Print("\n")
}


