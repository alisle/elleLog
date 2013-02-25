package Messages

import (
    "errors"
    "elleLog/elle/listener"
)

func newOSSIM(packet Listener.Packet) (*Message, error) {


    return &Message{}, errors.New("Message is not a valid AV packet")
}
