package Messages

import (
    "fmt"
    "log"
    "elleLog/elle/listener"
    "elleLog/elle/config"
)

type Message struct {
	Facility Facility 
	Severity Severity 
	TimeStamp string
	Hostname string
	Content string
    IP string
}

//External Globals
var MessagesReceived = 0
var LifetimeMessagesReceived = 0
var MAXTHREADS = 10

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

func (message *Message) String() string {
	return fmt.Sprintf("%s %v.%v %s %s", message.TimeStamp, message.Facility.String(), message.Severity.String(), message.Hostname,  message.Content)
}
func Initialize() {
    MAXTHREADS =  Config.GlobalConfig.GetInt(Config.MESSAGE_THREADS, 10)
}



func StartProcessing(finish chan bool, lines <-chan Listener.Packet, messages chan <- *Message) {
    for x := 0; x < MAXTHREADS; x++ {
        go MakeMessages(finish, lines, messages)
    }
}
func MakeMessages(finish chan bool, lines <-chan Listener.Packet, messages chan<- *Message) {

	for {
		select {
		case <- finish:
			{
				log.Print("Message: Signalled to end, closing")
                finish <- true
				return
			}
        case line := <- lines:
            {
                var newMessage *Message
                var err error

                if line.Type == Listener.RFC3164Packet {
                    newMessage, err = newRFC3164(line)
                } else {
                    newMessage, err = newOSSIM(line)
                }


                if err  == nil { 
                    if newMessage != nil  {
                        messages <- newMessage
                        MessagesReceived++
                    }
                } else {
                    log.Print("Messages:", err, " caused by ", line)
                }
            }
		}
	}
}

