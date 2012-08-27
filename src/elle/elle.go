package ellelog

import (
	"log"
	"bufio"
	"elle/rfc3164"
	"elle/listener"
	"elle/writers/logwriter"
	"elle/writers/stdoutwriter"
	"os"
	"os/signal"
	"regexp"
)

func ReadConf() {
	confRegex := regexp.MustCompile("^(?P<section>[^\\.]+)\\.(?P<var>[^=]+)=(?P<value>.*)")
	file, err := os.Open("etc/ellelog.cfg")
	if err != nil {
		log.Fatal("Error reading config file, aborting...")
	}

	reader := bufio.NewReader(file)

	line, _, err := reader.ReadLine()
	for err == nil {
		confLine := string(line)
		if len(confLine) > 2 {
			if matches := confRegex.FindStringSubmatch(confLine); matches != nil {
				log.Print("Found Config: ", confLine)
			} else {
				log.Print("Error reading config line: ", confLine)
			}
		}

		line, _, err = reader.ReadLine()
	}
}

func Run() {

	exit := make(chan bool)
	finished := make(chan bool, 3)
	lines := make(chan string, 1000)
	messages := make(chan *RFC3164.Message, 1000)

	ReadConf()
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
				StdoutWriter.Process(message)

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


