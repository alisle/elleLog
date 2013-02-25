package elleLog

// Imports
import (
	"log"
    "elleLog/elle/messages"
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
func Setup( finished chan bool, packets chan Listener.Packet, messages chan *Messages.Message ) {

    // Initialize the various components.
    ESWriter.Initialize()
    Messages.Initialize()

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
           SSWriter.New(server)
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

    if avLogger := Config.GlobalConfig.GetAllStrings(Config.LISTENER_ATTACH_AV_LOGGER); avLogger != nil {
        for _, logger := range avLogger {
            go Listener.AVListener(logger, finished, packets)
        }
    }

    
    Processors.LoadAllPlugins(Config.DEFAULT_PLUGIN_DIR)

}


func Run() {

	os.Chdir("../..")
	Config.WorkingDirectory, _ = os.Getwd()
    elleConfig, err := Config.New(Config.DEFAULT_CONFIG_FILE)
    if err != nil {
        log.Fatal("Unable to load Config: ", err)
    }

    Config.GlobalConfig = elleConfig


    exit := make(chan bool)
	finished := make(chan bool, 3)

    packets := make(chan Listener.Packet, elleConfig.GetInt(Config.MAX_QUEUE_PACKETS, 1000000))
	messages := make(chan *Messages.Message, elleConfig.GetInt(Config.MAX_QUEUE_MESSAGES, 1000000))
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

    Messages.StartProcessing(finished, packets, messages)

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
                    Messages.LifetimeMessagesReceived += Messages.MessagesReceived
                    Processors.LifetimeEventsReceived += Processors.EventsReceived

                    log.Print("Summary:")
                    log.Print("\t Packets/sec = ", Listener.PacketsReceived / summaryTime, " Events/sec = ", Processors.EventsReceived / summaryTime, " Messages/sec = ", Messages.MessagesReceived / summaryTime)
                    log.Print("\t Packet Queue = ", len(packets), " Event Queue = ", len(events), " Message Queue = ", len(messages))
                    log.Print("\t Life time Packets = ", Listener.LifetimePacketsReceived, " Life time Events = ", Processors.LifetimeEventsReceived, " Life time Messages = ", Messages.LifetimeMessagesReceived )
                    log.Print("\t Current number of GoRoutines: ", runtime.NumGoroutine())
                    Listener.PacketsReceived = 0
                    Messages.MessagesReceived = 0
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


