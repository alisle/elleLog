package RFC3164 
/* This processes lines in the RFC3164 format */

import (
	"errors"
	"regexp"
	"strconv"
	"log"
	"strings"
	"fmt"
    "elle/listener"
)

type Message struct {
	Facility Facility 
	Severity Severity 
	TimeStamp string
	Hostname string
	Content string
    IP string
}


type Facility int64
const (
	KERNEL Facility  = iota
	USER
	MAIL
	SYSTEM
	AUTH
	SYSLOG
	LPD
	NNTP
	UUCP
	CLOCK
	SECURITY
	FTP
	NTP
	CLOCK2
	LOCAL1
	LOCAL2
	LOCAL3
	LOCAL4
	LOCAL5
	LOCAL6
	LOCAL7
)

func (facility* Facility)  String() string {
	switch *facility {
		case KERNEL: { return "KERNEL"; }
		case USER: { return "USER"; }
		case MAIL: { return "MAIL"; }
		case SYSTEM: { return  "SYSTEM"; }
		case AUTH: { return "AUTH"; }
		case SYSLOG: { return "SYSLOG"; }
		case LPD: { return "LDP"; }
		case NNTP: { return "NNTP"; }
		case UUCP: { return "UUCP"; }
		case CLOCK: { return "CLOCK"; }
		case SECURITY:{ return "SECURITY"; }
		case FTP:{ return "FTP"; } 
		case NTP:{ return "NTP"; } 
		case CLOCK2:{ return "CLOCK2"; }
		case LOCAL1:{ return "LOCAL1"; }
		case LOCAL2:{ return "LOCAL2"; }
		case LOCAL3:{ return "LOCAL3"; }
		case LOCAL4:{ return "LOCAL4"; }
		case LOCAL5:{ return "LOCAL5"; }
		case LOCAL6:{ return "LOCAL6"; }
		case LOCAL7:{ return "LOCAL7"; }
	}

	return "nil";
}
type Severity int64

const (
	Emergency Severity = iota
	Alert
	Critical
	Error
	Warning
	Notice
	Information
	Debug
)

func (severity* Severity)  String() string {
	switch *severity {
	case Emergency : { return "Emergency";}
	case Alert : { return "Alert";}
	case Critical : { return "Critical";}
	case Error : { return "Error";}
	case Warning : { return "Warning";}
	case Notice : { return "Notice";}
	case Information : { return "Information";}
	case Debug : { return "Debug"; }
	}

	return "Invalid";
}
var messageRegex = regexp.MustCompile(`^<(?P<pri>\d+)>(?P<timestamp>\w{3}\s{1,2}\d{1,2}\s\d{2}:\d{2}:\d{2})\s(?P<Hostname>\S+)\s(?P<message>.*)`)

func New(packet Listener.Packet) (*Message, error) {
    line := packet.Message
	if matches := messageRegex.FindStringSubmatch(line); matches != nil {
		var facility, severity,pri int64
		var hostname string
		pri, _ = strconv.ParseInt(matches[1], 10, 8)

		facility = pri / 8
		severity = pri - (facility * 8) 
		timestamp := matches[2]

    hostname = strings.TrimSpace(matches[3])
	message := matches[4]

		return &Message{ Facility(facility), Severity(severity), timestamp, hostname, message, packet.Host}, nil
	}

	return &Message{}, errors.New("Message is not a valid RFC3164 packet")
}
func (message *Message) String() string {
	return fmt.Sprintf("%s %v.%v %s %s", message.TimeStamp, message.Facility.String(), message.Severity.String(), message.Hostname,  message.Content)
}

func MakeMessages(finish <-chan bool, lines <-chan Listener.Packet, messages chan<- *Message) {
	for {
		select {
		case <- finish:
			{
				log.Print("RFC3164: Signalled to end, closing")
				return
			}
		case line  := <-lines:
				newMessage, err := New(line)
				if err  == nil { 
					messages <- newMessage
				} else {
					log.Print("RFC3164:", err, " caused by ", line)
				}

		}
	}
}

