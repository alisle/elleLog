package RFC3164 
/* This processes lines in the RFC3164 format */

import (
	"errors"
	"regexp"
	"strconv"
	"log"
	"strings"
)

type Message struct {
	Facility int64
	Severity int64
	TimeStamp string
	Hostname string
	Tag string
	Content string
}

//var messageRegex = regexp.MustCompile(`^<(?P<facility>\d{1,2})(?P<severity>\d)>(?P<timestamp>\w{3}\s{1,2}\d{1,2}\s\d{2}:\d{2}:\d{2})\s(?P<hostname>\S+)\s(?P<tag>[^\[]+).*?:\s(?P<content>.*)`)

var messageRegex = regexp.MustCompile(`^<(?P<facility>\d{1,2})(?P<severity>\d)>(?P<timestamp>\w{3}\s{1,2}\d{1,2}\s\d{2}:\d{2}:\d{2})\s(?P<HostnameTag>[^:]+):(?P<message>.*)`)

func New(packet string) (*Message, error) {
	log.Print(packet)
	if matches := messageRegex.FindStringSubmatch(packet); matches != nil {
		var facility, severity int64
		var hostname, tag string

		facility, _ = strconv.ParseInt(matches[1], 10, 8)
		severity, _ = strconv.ParseInt(matches[2], 10, 8)
		timestamp := matches[3]

		hostTag := strings.TrimSpace(matches[4])
		if strings.Contains(hostTag, " ") {
			split :=  strings.Split(hostTag, " ")
			hostname = split[0]
			tag = split[1]
		} else {
			hostname = "127.0.0.1"
			tag = hostTag
		}
		message := matches[5]

		return &Message{ facility, severity, timestamp, hostname, tag, message}, nil
	}
	return &Message{}, errors.New("Message is not a valid RFC3164 packet")
}

func Process(finish <-chan bool, lines <-chan string, messages chan<- *Message) {
	for {
		select {
		case <- finish:
				log.Print("RFC3164: Signalled to end, closing")
				return
		case line  := <-lines:
				newMessage, err := New(line)
				if err  != nil { log.Print("RFC3164:", err); break; }
				messages <- newMessage
		}
	}
}

