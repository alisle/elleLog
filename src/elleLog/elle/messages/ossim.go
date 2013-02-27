package Messages

import (
    "elleLog/elle/listener"
    "strings"
)

func newOSSIM(packet Listener.Packet) (*Message, error) {
    line := packet.Message
    if strings.HasPrefix(line, "event") {
        return &Message{ SYSLOG, Information, "", "", line, packet.Host }, nil
    }

    return nil, nil 
}
