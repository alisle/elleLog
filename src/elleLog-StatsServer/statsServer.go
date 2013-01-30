package main

import (
	"elleLog-StatsServer/stats/connections"
    "elleLog-StatsServer/ws"
	"elleLog/elle/config"
	"elleLog/elle/processors"
	"encoding/gob"
	"log"
	"net"
	"os"
)

var events = make(chan Processors.Event, 100)

func handleConnection(conn net.Conn) {
	decoder := gob.NewDecoder(conn)
	count := 0
	for {
		count++
		var event Processors.Event
		if err := decoder.Decode(&event); err != nil {
			log.Print("Error when recieving: ", err)
			conn.Close()
			return
		}

        events <- event
	}
}

func launchServer() {
    go func() {
        service := Config.GlobalConfig.GetString(Config.SERVER_TCP_ADDRESS, ":4040")
        tcpAddr, err := net.ResolveTCPAddr("tcp", service)
        if err != nil {
            log.Fatal("Unable to Resolve Address: ", err)
        }

        listener, err := net.ListenTCP("tcp", tcpAddr)
        if err != nil {
            log.Fatal("Unable to Listen at Socket: ", err)
        }

        log.Print("Listening For elleLog Agents on ", service)
        for {
            conn, err := listener.Accept()
            if err != nil {
                log.Print("Problem with accept: ", err)
                continue
            }

            log.Print("elleLog Agent Connection Accepted: ", conn.RemoteAddr())

            go handleConnection(conn)

        }
    }()
}

func setup() {
	os.Chdir("../../")
	Config.WorkingDirectory, _ = os.Getwd()

	config, err := Config.New("etc/ellelog-statsserver.cfg")
	if err != nil {
		log.Fatal("Unable to load Config: ", err)
	}
	Config.GlobalConfig = config

	connections.Initalize()
    ws.Initialize()
}

func main() {
    setup()
	launchServer()
    ws.Launch()
    for event := range events {
        connections.Process(event)
        ws.Process(event)
    }
}
