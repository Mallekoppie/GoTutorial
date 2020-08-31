package CustomLogger

import (
	"log"
	"os"
)

var (
	Warn  *log.Logger
	Info  *log.Logger
	Error *log.Logger
)

func init() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	Info = log.New(file, "LogTesting INFO: ", log.Ldate|log.Ltime|log.Lshortfile|log.Lmsgprefix)
	Warn = log.New(file, "LogTesting WARNING: ", log.Ldate|log.Ltime|log.Lshortfile|log.Lmsgprefix)
	Error = log.New(file, "LogTesting ERROR: ", log.Ldate|log.Ltime|log.Lshortfile|log.Lmsgprefix)
}
