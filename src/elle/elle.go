package ellelog

import (
	"fmt"
	"elle/RFC3164"
	"elle/Listener"
)

func Run() {
	finished := make(chan bool)
	lines := make(chan string, 1000)
	messages := make(chan *RFC3164.Message, 1000)

	defer func() {
		close(finished)
		close(lines)
		close(messages)
	}()

	go Listener.UnixStreamListener("/dev/log", finished, lines)
	go Listener.UDPListener(":514", finished, lines)
	go RFC3164.Process(finished, lines, messages)

	for message := range messages {
		fmt.Println(message)
	}
}


