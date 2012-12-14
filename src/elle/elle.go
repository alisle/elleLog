package ellelog

import (
	"log"
	"elle/rfc3164"
	"elle/listener"
	"elle/writers/logwriter"
	"elle/writers/stdoutwriter"
	"elle/writers/eswriter"
	"elle/config"
	"elle/processors"
	"os"
	"os/signal"
	"strings"
)

func Setup( finished chan bool, packets chan Listener.Packet, messages chan *RFC3164.Message ) {
	elleConfig, err := Config.New("etc/ellelog.cfg")
	if err != nil {
		log.Fatal("Problem: ", err)
	}

	if showoutput, ok := elleConfig.GetVariable("output", "showstdout"); ok {
		if  strings.Contains(showoutput[0], "true") {
			log.Print("Outputing messages to STDOUT")
			StdoutWriter.ShowOutput = true
		}
	}

	if attachFiles, ok := elleConfig.GetVariable("output", "attachfile"); ok {
		for _, file := range attachFiles {
			log.Print("Attaching to log file: ", file)
			LogWriter.New(file)
		}
	}

	if attachES, ok := elleConfig.GetVariable("output", "attachelastisearch"); ok {
		for _, file := range attachES {
			log.Print("Attaching ElastiSearch Server: ", file)
			ESWriter.New(file)
		}
	}

	if attachUDP, ok := elleConfig.GetVariable("listener", "attachudp"); ok {
		for _, host := range attachUDP {
			log.Print("Starting UPD Listener on ", host)
			go Listener.UDPListener(host, finished, packets)
		}
	}

	if attachUS, ok := elleConfig.GetVariable("listener", "attachunixstream"); ok {
		for  _, unixstream := range attachUS {
			log.Print("Starting UnixStreamListener")
			go Listener.UnixStreamListener(unixstream, finished, packets)
		}
	}

	Processors.LoadAllPlugins("etc/plugins/")
}

func Run() {

	os.Chdir("../..")
	Config.WorkingDirectory, _ = os.Getwd()

	exit := make(chan bool)
	finished := make(chan bool, 3)
	packets := make(chan Listener.Packet, 1000)
	messages := make(chan *RFC3164.Message, 1000)
	defer func() {
		close(packets)
		close(messages)
		close(finished)
		LogWriter.Close()
		ESWriter.Close()
	}()

	Setup(finished, packets, messages)
	log.Print("Launching Raw Log to RFC3164 Message")
	go RFC3164.MakeMessages(finished, packets, messages)


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
                Processors.CheckMessage(message)
				//LogWriter.Process(message)
				//StdoutWriter.Process(message)
				//ESWriter.Process(message)

			case <- exit:
				log.Print("Caught Sigterm")
				for x := 0; x < 3; x++ {
					finished <- true
				}
				return
		}
	}

}


