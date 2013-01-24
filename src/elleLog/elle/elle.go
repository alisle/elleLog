package elleLog

// Imports
import (
	"log"
    "elleLog/elle/rfc3164"
    "elleLog/elle/listener"
    "elleLog/elle/writers/logwriter"
    "elleLog/elle/writers/stdoutwriter"
    "elleLog/elle/writers/eswriter"
    "elleLog/elle/writers/sswriter"
    "elleLog/elle/config"
    "elleLog/elle/processors"
	"os"
	"os/signal"
   "time"
   "runtime"
)

// External Globals
func Setup( finished chan bool, packets chan Listener.Packet, messages chan *RFC3164.Message ) {
    ESWriter.Initialize()
    RFC3164.Initialize()

    StdoutWriter.ShowOutput = Config.GlobalConfig.GetBool(Config.OUTPUT_SHOWSTDOUT, false)
    if files := Config.GlobalConfig.GetAllStrings(Config.OUTPUT_ATTACH_FILE); files != nil {
        for _, file := range files {
            LogWriter.New(file)
        }
    }

    if elastics := Config.GlobalConfig.GetAllStrings(Config.OUTPUT_ATTACH_ELASTISEARCH); elastics !=  nil {
        for _, elastic := range elastics {
            ESWriter.New(elastic)
        }
    }


    if stats := Config.GlobalConfig.GetAllStrings(Config.OUTPUT_ATTACH_STATSERVER); stats != nil {
        for _, server := range stats {
            go SSWriter.New(server)
        }
    } 

    if udp := Config.GlobalConfig.GetAllStrings(Config.LISTENER_ATTACH_UDP); udp != nil {
        for _, port := range udp {
            go Listener.UDPListener(port, finished, packets)
        }
    }

    if unixStream := Config.GlobalConfig.GetAllStrings(Config.LISTENER_ATTACH_UNIX_STREAM); unixStream != nil {
        for _, file := range unixStream {
            go Listener.UnixStreamListener(file, finished, packets)
        }
    }
    Processors.LoadAllPlugins("etc/plugins/")

}

func Run() {

	os.Chdir("../..")
	Config.WorkingDirectory, _ = os.Getwd()
    elleConfig, err := Config.New("etc/ellelog.cfg")
    if err != nil {
        log.Fatal("Unable to load Config: ", err)
    }

    Config.GlobalConfig = elleConfig


    exit := make(chan bool)
	finished := make(chan bool, 3)
	packets := make(chan Listener.Packet,   elleConfig.GetInt(Config.MAX_QUEUE_PACKETS, 1000000))
	messages := make(chan *RFC3164.Message, elleConfig.GetInt(Config.MAX_QUEUE_MESSAGES, 1000000))
    events := make(chan Processors.Event,   elleConfig.GetInt(Config.MAX_QUEUE_EVENTS, 1000000))

	defer func() {
		close(packets)
		close(messages)
		close(finished)
		LogWriter.Close()
		ESWriter.Close()
	}()

    log.Print("Maximum number of CPUS: ", runtime.NumCPU())
    log.Print("Setting maximum number of CPUS: ", elleConfig.GetInt(Config.MAX_CPUS, 1))
    runtime.GOMAXPROCS(elleConfig.GetInt(Config.MAX_CPUS, 1))


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
    summaryTime := elleConfig.GetInt(Config.MAX_SUMMARY_TIME, 30)
    go func() {
        for {
            select {
                case <- time.After(time.Duration(summaryTime) * time.Second): 
                {
                    Listener.LifetimePacketsReceived += Listener.PacketsReceived
                    RFC3164.LifetimeMessagesReceived += RFC3164.MessagesReceived
                    Processors.LifetimeEventsReceived += Processors.EventsReceived

                    log.Print("Summary:")
                    log.Print("\t Packets/sec = ", Listener.PacketsReceived / summaryTime, " Events/sec = ", Processors.EventsReceived / summaryTime, " Messages/sec = ", RFC3164.MessagesReceived / summaryTime)
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
            case event := <- events:
                StdoutWriter.Process(event)
                ESWriter.Process(event)
                SSWriter.Process(event)


            case <- exit:
                log.Print("Caught Sigterm")
                for x := 0; x < 3; x++ {
                    finished <- true
                }
                return
        }
    }

}


