package ESWriter


// Imports
import (
	"log"
	"encoding/json"
	"bytes"
	"net/http"
	"io/ioutil"
    "elleLog/elle/processors"
    "time"
)

// External Globals
var BULK_MAX_ITEMS = 50
var BULK_USE = true
var BULK_SECONDS time.Duration = 1
var MAX_CONNECTIONS = 5

// Internal Globals
var servers =  make([]*ESWriter, 10)
var currentServer = 0

// Types
type okPacket struct {
	OK bool
	Index string
	Type string
	Id string
	Version int
}

type ESWriter struct {
	URL string 
    events chan Processors.Event
}
func (esWriter *ESWriter)postSingleHTTP(event Processors.Event) {
    jsonPacket, err := json.Marshal(event)
    if err != nil {
        log.Print("Unable to Marshal Message: ", err)
        return
    } else {

        response, err := esWriter.post(bytes.NewBuffer(jsonPacket), "") 
        if err != nil {
            return
        }

        var  returnValue okPacket
        err = json.Unmarshal(response, &returnValue)
        if err != nil {
            log.Print("Unable to verify elastisearch received packet!", err)

            return
        }

        if returnValue.OK != true {
            log.Print("Filed to add to server:", esWriter.URL, " reason: ", string(response))
        }
    }
}

func (esWriter *ESWriter)post(buffer *bytes.Buffer, tags string) ([]byte, error ){
    resp, err := http.Post(esWriter.URL + tags, "applicaton/json", buffer)
    if err != nil {
        log.Print("Unable to bulk send to server: ", err, " server:", esWriter.URL)
        return nil, err
    }

    defer resp.Body.Close()
    response, _ := ioutil.ReadAll(resp.Body)

    return response, nil

}

func (esWriter *ESWriter)singleProcess() {
    for message := range esWriter.events {
        esWriter.postSingleHTTP(message)
    }
}

func (esWriter *ESWriter)bulkProcess() {
   for {
       var currentItems = 0
       var buffer bytes.Buffer
        
       ITEM_LOOP:
       for currentItems < BULK_MAX_ITEMS {
           select {
                case <- time.After(BULK_SECONDS * time.Second):
                   if currentItems > 0 {
                       break ITEM_LOOP
                   }
               case event := <- esWriter.events:
                   jsonPacket, err := json.Marshal(event)
                   if err != nil {
                       log.Print("Unable to Marshal Message: ", err)
                       continue;
                   } else {
                       buffer.WriteString("{\"index\": {}}\n")
                       buffer.Write(jsonPacket)
                       buffer.WriteString("\n")
                       currentItems++
                   }
                  
           }
       }

        _, err := esWriter.post(&buffer, "/_bulk")
        if err != nil {
            continue
        }

        if err != nil {
            log.Print("Unable to verify elasticsearch received packet!", err)
            return
        }

    }

}
func (esWriter *ESWriter)Process() {
    for x := 0; x < MAX_CONNECTIONS; x++ {
        if BULK_USE {
                go esWriter.bulkProcess()
            } else  {
                go esWriter.singleProcess()
            }
    }
    
}


func New(host string) {
	if currentServer > 9 {
		log.Print("Too many ElastiSearch Writers!")
	} else {
		log.Print("Attached new ElastiSearch output:", host)
        events := make(chan Processors.Event, 1000)
		esWriter := &ESWriter{host,  events}
		servers[currentServer] = esWriter
		currentServer += 1
		esWriter.Process()
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
