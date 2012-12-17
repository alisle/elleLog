package ellelog

// Imports
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
   "time"
   "runtime"
)
// External Globals
var NUMCPU = 4

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
	packets := make(chan Listener.Packet,   1000000)
	messages := make(chan *RFC3164.Message, 1000000)
    events := make(chan Processors.Event,   1000000)
	defer func() {
		close(packets)
		close(messages)
		close(finished)
		LogWriter.Close()
		ESWriter.Close()
	}()
    runtime.GOMAXPROCS(NUMCPU)

    log.Print("Maximum number of CPUS: ", runtime.NumCPU())
    log.Print("Setting number of CPUS: ",  runtime.GOMAXPROCS(NUMCPU))


	Setup(finished, packets, messages)
	log.Print("Launching Raw Log to RFC3164 Message")

    RFC3164.StartProcessing(finished, packets, messages)

    Processors.AttachMsgChannel(messages)
    Processors.AttachEventsChannel(events)

    Processors.StartProcessing() 

	catchSIGTERM := make(chan os.Signal, 1)
	signal.Notify(catchSIGTERM, os.Interrupt)
	go func() {
		for _ = range catchSIGTERM {
			exit <- true
			return
		}
	}()

    go func() {
        for {
            select {
                case <- time.After(2 * time.Second): 
                {
                    Listener.LifetimePacketsReceived += Listener.PacketsReceived
                    RFC3164.LifetimeMessagesReceived += RFC3164.MessagesReceived
                    Processors.LifetimeEventsReceived += Processors.EventsReceived

                    log.Print("Sumary:")
                    log.Print("\t Packets/sec = ", Listener.PacketsReceived / 2, " Events/sec = ", Processors.EventsReceived / 2, " Messages/sec = ", RFC3164.MessagesReceived / 2)
                    log.Print("\t Packet Queue = ", len(packets), " Event Queue = ", len(events), " Message Queue = ", len(messages))
                    log.Print("\t Life time Packets = ", Listener.LifetimePacketsReceived, " Life time Events = ", Processors.LifetimeEventsReceived, " Life time Messages = ", RFC3164.LifetimeMessagesReceived )
                    log.Print("\t Current number of GoRoutines: ", runtime.NumGoroutine())
                    Listener.PacketsReceived = 0
                    RFC3164.MessagesReceived = 0
                    Processors.EventsReceived = 0
                }
            }
        }
    }()


	for
	{
		select {
            case _ = <- events:
                //StdoutWriter.Process(event)
                //ESWriter.Process(event)

				//LogWriter.Process(message)
				//StdoutWriter.Process(message)

			case <- exit:
				log.Print("Caught Sigterm")
				for x := 0; x < 3; x++ {
					finished <- true
				}
				return
		}
	}

}


