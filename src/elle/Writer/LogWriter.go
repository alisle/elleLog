package Writer

import ( 
	"elle/RFC3164"
	"os"
	"bufio"
)


type struct LogWriter {
	file *File,
	writer *Writer
	FileName,
}

func NewLogWriter(fileName string) *LogWriter, error {
	file, err :=  os.OpenFile(fileName, O_APPEND,  0666)
	if err == nil
		writer := bufio.NewWriter(file)

	return &LogWriter{file, writer, fileName}, err
}


func (LogWriter* logWriter)WriteString(string  output) {
	logWriter.WriteString(
}
