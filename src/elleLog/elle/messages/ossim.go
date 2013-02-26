package Messages

import (
    "elleLog/elle/listener"
)

func newOSSIM(packet Listener.Packet) (*Message, error) {
    line := packet.Message
    return &Message{ SYSLOG, Information, "", "", line, packet.Host }, nil
}
