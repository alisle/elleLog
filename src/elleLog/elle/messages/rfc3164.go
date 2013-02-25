package Messages
/* This processes lines in the RFC3164 format */

// Imports
import (
	"errors"
	"regexp"
	"strconv"
	"strings"
    "elleLog/elle/listener"
)


//Internal Globals
var syslogRegex = regexp.MustCompile(`^<(?P<pri>\d+)>(?P<timestamp>\w{3}\s{1,2}\d{1,2}\s\d{2}:\d{2}:\d{2})\s(?P<Hostname>\S+)\s(?P<message>.*)`)

func newRFC3164(packet Listener.Packet) (*Message, error) {
    line := packet.Message
	if matches := syslogRegex.FindStringSubmatch(line); matches != nil {
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


