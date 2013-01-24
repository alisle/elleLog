package SSWriter


// Imports
import (
	"log"
	"encoding/gob"
    "net"
    "elleLog/elle/processors"
)

// Internal Globals
var servers =  make([]*SSWriter, 10)
var currentServer = 0

type SSWriter struct {
	Addr string 
    events chan Processors.Event
    conn net.Conn
    encoder *gob.Encoder
}

func Initialize() {
}

func (ssWriter *SSWriter)Process() {
    for event := range ssWriter.events {
        err := ssWriter.encoder.Encode(event)
        if err != nil {
            log.Print("Unable to send event to StatServer ", ssWriter.Addr, " returning")
            return
        }
    }
}


func New(host string) {
	if currentServer > 9 {
		log.Print("Too many StatServer Writers!")
	} else {
        conn, err := net.Dial("tcp", host)
        if err != nil {
            log.Print("Unable to attach: ", host)
            return
        }
		log.Print("Attached new StatServer output:", host)
        events := make(chan Processors.Event, 1000)
		ssWriter := &SSWriter{host,  events, conn, gob.NewEncoder(conn)}
		servers[currentServer] = ssWriter
		currentServer += 1
		ssWriter.Process()
	}
}

func Process(event Processors.Event) {
    for x:= 0; x < currentServer; x++ {
        servers[x].events <- event
    }
}

func Close() {
	for x:= 0; x < currentServer; x++ {
		close(servers[x].events)
	}
}
