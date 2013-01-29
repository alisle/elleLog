package SSWriter


// Imports
import (
    "errors"
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

func (ssWriter *SSWriter)Connect() error {
        conn, err := net.Dial("tcp", ssWriter.Addr)
        if err != nil {
            log.Print("Unable to attach: ", ssWriter.Addr)
            return errors.New("Failed")
        }

        ssWriter.conn = conn
        ssWriter.encoder = gob.NewEncoder(ssWriter.conn)
        return nil
}
func (ssWriter *SSWriter)Process() {
    for event := range ssWriter.events {
        err := ssWriter.encoder.Encode(event)
        if err != nil {
            log.Print("Unable to send event to StatServer ", ssWriter.Addr, " reconnecting")
            if err = ssWriter.Connect(); err != nil {
                return
            }
        }
    }
}


func New(host string) {
	if currentServer > 9 {
		log.Print("Too many StatServer Writers!")
	} else {
        log.Print("Attached new StatServer output:", host)
		ssWriter := &SSWriter{host,  make(chan Processors.Event, 1000), nil, nil}
        if err := ssWriter.Connect(); err == nil {
            servers[currentServer] = ssWriter
            currentServer += 1
            go ssWriter.Process()
        }
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
