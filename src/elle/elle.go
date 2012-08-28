package ellelog

import (
	"log"
	"bufio"
	"elle/rfc3164"
	"elle/listener"
	"elle/writers/logwriter"
	"elle/writers/stdoutwriter"
	"elle/writers/eswriter"
	"os"
	"os/signal"
	"regexp"
	"strings"
)

func ReadConf( finished chan bool, lines chan string, messages chan *RFC3164.Message ) {
	confRegex := regexp.MustCompile("^(?P<section>[^\\.]+)\\.(?P<var>[^=]+)=(?P<value>.*)")
	os.Chdir("../..")
	workingDir, _ := os.Getwd()
	configFile := workingDir + "/etc/ellelog.cfg"
	log.Print("Loading Config: ", configFile)
	file, err := os.Open(configFile)
	if err != nil {
		log.Fatal("Error reading config file, aborting...")
	}

	reader := bufio.NewReader(file)

	line, _, err := reader.ReadLine()
	for err == nil {
		confLine := string(line)
		if len(confLine) > 2 {
			if matches := confRegex.FindStringSubmatch(confLine); matches != nil {
				section := strings.ToLower(matches[1])
				variable := strings.ToLower(matches[2])
				value := strings.Trim(matches[3], "\" ")

				switch(section) {
					case "output": {
						switch(variable) {
							case "showstdout": { 
								if  strings.Contains(value, "true") {
									log.Print("Outputing messages to STDOUT")
									StdoutWriter.ShowOutput = true
								}
							}
							case "attachfile": { 
								log.Print("Attaching to log file: ", value)
								LogWriter.New(value)
							}
							case "attachelastisearch": {
								log.Print("Attaching ElastiSearch Server: ", value)
								ESWriter.New(value) 
							}
						}
					}
					case "listener": {
						switch(variable) {
							case "attachudp": {
								log.Print("Starting UPD Listener on ", value)
								go Listener.UDPListener(value, finished, lines)
							}

							case "attachunixstream": {
								log.Print("Starting UnixStreamListener")
								go Listener.UnixStreamListener(value, finished, lines)
							}
						}
					}
				}

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
	defer func() {
		close(lines)
		close(messages)
		close(finished)
		LogWriter.Close()
		ESWriter.Close()
	}()

	ReadConf(finished, lines, messages)
	log.Print("Launching Raw Log to RFC3164 Message")
	go RFC3164.MakeMessages(finished, lines, messages)


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
				ESWriter.Process(message)

			case <- exit:
				log.Print("Caught Sigterm")
				for x := 0; x < 3; x++ {
					finished <- true
				}
				return
		}
	}

}


