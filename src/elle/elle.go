package ellelog

import (
	"log"
	"elle/rfc3164"
	"elle/listener"
	"elle/writers/logwriter"
	"os"
	"os/signal"
)

func Run() {
	exit := make(chan bool)
	finished := make(chan bool, 3)
	lines := make(chan string, 1000)
	messages := make(chan *RFC3164.Message, 1000)

	logoutput, _ := LogWriter.New("temp.log")

	defer func() {
		close(lines)
		close(messages)
		close(finished)
		logoutput.Close()
	}()

	log.Print("Starting UnixStreamListener")
	go Listener.UnixStreamListener("/dev/log", finished, lines)
	log.Print("Starting UPD Listener")
	go Listener.UDPListener(":514", finished, lines)
	log.Print("Launching Raw Log to RFC3164 Message")
	go RFC3164.MakeMessages(finished, lines, messages)
	log.Print("Starting Message Processors")


	catchSIGTERM := make(chan os.Signal, 1)
	signal.Notify(catchSIGTERM, os.Interrupt)
	go func() {

		for _ = range catchSIGTERM {
			exit <- true
			return
		}
	}()

	for
	{
		select {
			case message := <- messages:
				LogWriter.Process(message)
			case <- exit:
				{
					log.Print("Caught Sigterm")
					for x := 0; x < 3; x++ {
						finished <- true
					}
					return
				}
		}
	}

}


