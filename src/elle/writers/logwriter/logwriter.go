package LogWriter

import ( 
	"log"
	"elle/rfc3164"
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
	logWriter.writer.WriteString(output + "\n")
	logWriter.writer.Flush()
}

func (logWriter *LogWriter)Close() {
	logWriter.file.Close()
}

var writers = make([]*LogWriter, 10)
var currentWriter = 0

func New(fileName string) (*LogWriter, error) {
	var logWriter *LogWriter
	file, err :=  os.OpenFile(fileName, os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0666)

	if err == nil {
		if currentWriter > 9 {
			log.Print("Too many log writers!")
		} else {
			log.Print("Attached new output:", fileName)
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

func Process(message *RFC3164.Message) {
	for x := 0; x < currentWriter; x ++ {
		writer := writers[x]
		writer.messages <- message
	}
}

