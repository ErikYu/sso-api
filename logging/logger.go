package logging

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
)

var logger *log.Logger
var logPrefix = ""
var levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
var defaultPrefix = ""

type level int

const (
	DEBUG level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func init() {
	logger = log.New(openLogFile(getLogFileFullPath()), defaultPrefix, log.LstdFlags)
}

func setPrefix(level level) {
	_, file, line, ok := runtime.Caller(2)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}
	logger.SetPrefix(logPrefix)
}
func Info(format string, v ...interface{}) {
	setPrefix(INFO)
	logger.Printf(format, v)
}
func Error(format string, v ...interface{}) {
	setPrefix(ERROR)
	logger.Printf(format, v)
}
