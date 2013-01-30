package connections

import (
    "log"
    "elleLog/elle/processors"
    "elleLog/elle/config"
    "time"
)


type IP struct {
    Addr string
    Destinations []IP 
    Sources []IP
}

func (this *IP)AddDestination(dst IP) {
    if this.containsIP(this.Destinations, dst) == false {
        this.Destinations = append(this.Destinations, dst)
    }
}

func (this *IP)AddSource(src IP) {
    if this.containsIP(this.Sources, src) == false {
        this.Sources = append(this.Sources, src)
    }
}

func (this *IP)containsIP(ips []IP, ip IP) bool {
    for _, current := range ips {
        if current.Addr == ip.Addr {
            return true
        }
    }

    return false
}
type Graph map[string] IP

func (this *Graph) AddConnection(_src, _dst IP) {
    src := (*this)[_src.Addr]
    src.AddDestination(_dst)

    dst := (*this)[_dst.Addr]
    dst.AddSource(_src)
}
var ConnectionGraph Graph

func Initalize() {
    ConnectionGraph = make(Graph)
    DisplaySummary()
}

func getIP(addr string) IP {
    var ip IP
    if val, ok := ConnectionGraph[addr]; ok {
        ip = val
    } else {
        ip = IP{addr, make([]IP, 0, 10), make([]IP, 0, 10)}
        ConnectionGraph[addr] = ip
    }

    return ip
}



func GrabConnections(event Processors.Event) {
    if src, ok := event["source_address"]; ok {
        if dst, ok := event["destination_address"]; ok {
            srcip := getIP(src)
            dstip := getIP(dst)
            ConnectionGraph.AddConnection(srcip, dstip)
        }
    }
}

func DisplaySummary() {
    summaryTime := Config.GlobalConfig.GetInt(Config.MAX_SUMMARY_TIME, 30)

    go func() {
        for {
            select { 
                case <- time.After(time.Duration(summaryTime) * time.Second): {
                    log.Print("Connection Summary:")
                    log.Print("\tNumber of Hosts: ", len(ConnectionGraph))
                }
            }
        }
    } ()
}

