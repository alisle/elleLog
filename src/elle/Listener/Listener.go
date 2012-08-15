package Listener

import (
	"net"
	"log"
	"os"
)



func UDPListener(port string, finish <-chan bool, lines chan<- string) {
	listener("udp", port, finish , lines)
}

func UnixDatagramListener(fileName string, finish <-chan bool, lines chan<- string) {
	listener("unixgram", fileName, finish, lines)
}

func UnixStreamListener(fileName string, finish <-chan bool, lines chan<- string) {

	var err error 
	if _, err := os.Stat(fileName); err == nil {
		log.Print(fileName, " exists deleting..")
		err = os.Remove(fileName)
		if err != nil {
			log.Print("Unable to delete file")
			return
		}
	}

	listen, err := net.Listen("unix", fileName)
	if err != nil {
		log.Print("Unable to Listen to file: ", err)
		return
	}

	err = os.Chmod(fileName, 0666)
	if err != nil {
		log.Print("Unable to chmod file: ", err)
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			continue
		}

		go func() {
			for  {
				buffer := make([]byte, 1024)
				bytesRead, err := conn.Read(buffer[0:])
				if err != nil {
					log.Print("Connection Closed.")
					 break;
				} else {
					 lines <- string(buffer[0:bytesRead])
				}
			}
		}()
	}
}

func listener(prot string, url string, finish <-chan bool, lines chan<- string) {
	listener, err := net.ListenPacket(prot, url)
	if err != nil { 
		log.Print("ListenPacket Failure: ", err, " not listening") 
		return
	}

	go func() {
		for  {
			buffer := make([]byte, 1024)
			bytesRead, _, err := listener.ReadFrom(buffer[0:])
			if err != nil {
				 log.Print("Listener: Unable to Read Packet!")
			} else {
				 lines <- string(buffer[0:bytesRead])
			}
		}
	}()

	for  {

		select  {
		case <- finish:
			log.Print("listener: signalled to end, closing")
			return
		}
	}
}
