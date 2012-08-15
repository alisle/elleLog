package Writer

import ( 
	"elle/RFC3164"
	"os"
)


type struct LogWriter {
	file *File,
	FileName,
}

func NewLogWriter(fileName string) *LogWriter, error {
	file, err =  os.OpenFile(fileName, O_APPEND,  0666)
	return &LogWriter{file, fileName}, err
}


