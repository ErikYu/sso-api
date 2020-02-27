package logging

import (
	"fmt"
	"log"
	"os"
)

var simpleLogger *log.Logger

func init() {
	simpleLogger = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
}

func STDDebug(format string, v ...interface{}) {
	simpleLogger.SetPrefix("[DEBUG]")
	_ = simpleLogger.Output(2, fmt.Sprintf(format, v...))
}

func STDInfo(format string, v ...interface{}) {
	simpleLogger.SetPrefix("[INFO]")
	_ = simpleLogger.Output(2, fmt.Sprintf(format, v...))
}

func STDError(format string, v ...interface{}) {
	simpleLogger.SetPrefix("[ERROR]")
	_ = simpleLogger.Output(2, fmt.Sprintf(format, v...))
}
