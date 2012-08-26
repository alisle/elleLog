package LogWriter

import ( 
	"fmt"
	"log"
	"elle/RFC3164"
	"os"
	"bufio"
)

type LogWriter struct {
	file *os.File
	writer *bufio.Writer
	messages chan *RFC3164.Message
	FileName string
}
func (logWriter *LogWriter)Process()  {
	for message := range logWriter.messages {
		logWriter.WriteMessage(message)
	}
}
func (logWriter *LogWriter)WriteMessage(msg *RFC3164.Message) {
	logWriter.WriteString(msg.String())
}

func (logWriter *LogWriter)WriteString(output string) {
	logWriter.file.WriteString(output + "\n")
}

func (logWriter *LogWriter)Close() {
	logWriter.file.Close()
}

var writers = make([]*LogWriter, 10)
var currentWriter = 0

func New(fileName string) (*LogWriter, error) {
	var logWriter *LogWriter
//	file, err :=  os.OpenFile(fileName, os.O_APPEND | os.O_CREATE, 0666)
	file, err := os.Create(fileName)

	if err == nil {
		if currentWriter > 9 {
			log.Print("Too many log writers!")
		} else {
			_, newErr := file.WriteString("Hello!")
			fmt.Println(newErr)
			writer := bufio.NewWriter(file)
			messages := make(chan *RFC3164.Message, 500)
			logWriter = &LogWriter{file, writer, messages, fileName}
			writers[currentWriter] = logWriter
			currentWriter += 1
			go logWriter.Process()

			return logWriter, err
		}
	}

	return  nil, err
}

func Process(finish <-chan bool, messages chan *RFC3164.Message) {
	go func() {
		for message := range messages {
			for x := 0; x < currentWriter; x ++ {
				writer := writers[x]
				writer.messages <- message
			}
		}
	}()

	for {
		select {
		case <- finish:
			log.Print("logwriter: signalled to end, closing")
			return
		}
	}
}

