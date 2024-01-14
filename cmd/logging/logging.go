package logging

import (
	"fmt"
	"io"
	"log"
)

type Logger interface {
	Info(v ...any)
	Error(v ...any)
	Debug(v ...any)
}

type GutHubLogger struct {
	infoLog  *log.Logger
	errLog   *log.Logger
	debugLog *log.Logger
}

func NewGutHubLogger(infoOut, debugOut, errOut io.Writer, prefix string, flag int) *GutHubLogger {
	return &GutHubLogger{
		infoLog:  log.New(infoOut, prefix+"INFO: ", flag),
		debugLog: log.New(debugOut, prefix+">>> DEBUG: ", flag),
		errLog:   log.New(errOut, prefix+"ERROR: ", flag),
	}
}

func (l *GutHubLogger) Info(v ...any) {
	l.infoLog.Print(fmt.Sprintln(v...))
}

func (l *GutHubLogger) Error(v ...any) {
	l.errLog.Print(fmt.Sprintln(v...))
}

func (l *GutHubLogger) Debug(v ...any) {
	l.debugLog.Print(fmt.Sprintln(v...))
}
