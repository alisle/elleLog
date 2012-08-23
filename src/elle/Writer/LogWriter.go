package LogWriter

import ( 
	"elle/RFC3164"
	"os"
	"bufio"
)

type LogWriter struct {
	file *os.File
	writer *bufio.Writer
	FileName string
}

func NewLogWriter(fileName string) (*LogWriter, error) {
	var writer  *bufio.Writer
	file, err :=  os.OpenFile(fileName, os.O_APPEND,  0666)
	
	if err == nil { writer = bufio.NewWriter(file) }

	return &LogWriter{file, writer, fileName}, err
}

func (logWriter *LogWriter)WriteMessage(msg RFC3164.Message) {
	logWriter.WriteString(msg.String())
}

func (logWriter *LogWriter)WriteString(output string) {
	logWriter.writer.WriteString(output)
}

func (logWriter *LogWriter)Close() {
	logWriter.file.Close()
}
