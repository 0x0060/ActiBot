package logger

import (
	"fmt"
	"log"
)

type LogLevel string

const (
	System LogLevel = "SYS"
	Info   LogLevel = "INF"
	Error  LogLevel = "ERR"
	Debug  LogLevel = "DBG"
	API    LogLevel = "API"
)

func Log(level LogLevel, message string) {
	log.Printf("(%s): %s", level, message)
}

func Logf(level LogLevel, format string, v ...interface{}) {
	log.Printf("(%s): %s", level, fmt.Sprintf(format, v...))
}
