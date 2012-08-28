package ESWriter

import (
	"log"
	"encoding/json"
	"bytes"
	"net/http"
	"elle/rfc3164"
	"io/ioutil"
)
type okPacket struct {
	OK bool
	Index string
	Type string
	Id string
	Version int
}

type ESWriter struct {
	URL string 
	messages chan *RFC3164.Message
}
func (esWriter *ESWriter)Process() {
	for message := range esWriter.messages  {
		jsonPacket, err := json.Marshal(message)
		if err != nil {
			log.Print("Unable to Marshal Message: ", err)
			continue;
		} else {
			resp, err := http.Post(esWriter.URL, "application/json", bytes.NewBuffer(jsonPacket))
			if err != nil {
				log.Print("Unable to send to server: ", err, " server:", esWriter.URL) 
				continue;
			}

			defer resp.Body.Close()
			var  returnValue okPacket
			body, _ := ioutil.ReadAll(resp.Body)
			err = json.Unmarshal(body, &returnValue)
			if err != nil {
				log.Print("Unable to verify elastisearch received packet!", err)
				continue;
			}

			if returnValue.OK != true {
				log.Print("Filed to add to server:", esWriter.URL, " reason: ", string(body))
			}
		}
	}
}
var servers =  make([]*ESWriter, 10)
var currentServer = 0

func New(host string) {
	if currentServer > 9 {
		log.Print("Too many ElastiSearch Writers!")
	} else {
		log.Print("Attached new ElastiSearch output:", host)
		messages := make(chan *RFC3164.Message, 1000)
		esWriter := &ESWriter{host, messages}
		servers[currentServer] = esWriter
		currentServer += 1
		go esWriter.Process()
	}
}

func Process(message *RFC3164.Message) {
	for x:= 0; x < currentServer; x++ {
		servers[x].messages  <- message
	}
}

func Close() {
	for x:= 0; x < currentServer; x++ {
		close(servers[x].messages)
	}
}
