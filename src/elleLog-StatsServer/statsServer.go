package main

import (
    "elleLog/elle/config"
    "elleLog/elle/processors"
    "os"
    "log"
    "net"
    "encoding/gob"
)

func handleConnection(conn net.Conn) {
    decoder := gob.NewDecoder(conn)

    for {
        var event Processors.Event
        if err := decoder.Decode(&event); err != nil {
            log.Print("Error when recieving: ", err)
            conn.Close()
            return
        }
        grabConnections(event)
    }
}

func grabConnections(event Processors.Event) {

    if src, ok := event["source_address"]; ok {
        if dst, ok := event["destination_address"]; ok {
            log.Print(src, " -> ", dst)
        }
    }

}
func main() {
    os.Chdir("../../")
    Config.WorkingDirectory, _ = os.Getwd()

    config, err := Config.New("etc/ellelog-statsserver.cfg")
    if err != nil {
        log.Fatal("Unable to load Config: ", err)
    }

    Config.GlobalConfig = config

    service := Config.GlobalConfig.GetString(Config.SERVER_TCP_ADDRESS, ":4040")
    tcpAddr, err := net.ResolveTCPAddr("tcp", service)
    if err != nil {
        log.Fatal("Unable to Resolve Address: ", err)
    }

    listener, err := net.ListenTCP("tcp", tcpAddr)
    if err != nil {
        log.Fatal("Unable to Listen at Socket: ", err)
    }

    log.Print("Listening on ", service)
    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Print("Problem with accept: ", err)
            continue
        }

        log.Print("Accepted Connection: ", conn)

        go handleConnection(conn)

    }


}
