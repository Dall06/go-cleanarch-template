package server

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

type LoggerHandler struct{}

// NewLogger return a Logger.
func NewLoggerHandler() *LoggerHandler {
	return &LoggerHandler{}
}

func (l *LoggerHandler) Open() {
	filePath := ".env"

	f, err := os.Open(filePath)
	if err != nil {
		l.LogError("%s", err)
	}
	defer f.Close()

	lines := make([]string, 0, 100)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		l.LogError("%s", err)
	}

	for _, l := range lines {
		pair := strings.Split(l, "=")
		os.Setenv(pair[0], pair[1])
	}
}

// LogError is print messages to log.
func (*LoggerHandler) LogError(format string, v ...interface{}) {
	file, err := os.OpenFile("log/error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Printf("%s", err)
	}
	defer file.Close()

	log.SetOutput(io.MultiWriter(file, os.Stdout))
	log.SetFlags(log.Ldate | log.Ltime)

	log.Printf(format, v...)
}

// LogAccess is print messages to log.
func (*LoggerHandler) LogAccess(format string, v ...interface{}) {
	file, err := os.OpenFile("log/access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Printf("%s", err)
	}
	defer file.Close()

	log.SetOutput(io.MultiWriter(file, os.Stdout))
	log.SetFlags(log.Ldate | log.Ltime)

	log.Printf(format, v...)
}
