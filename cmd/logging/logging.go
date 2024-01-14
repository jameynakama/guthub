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
	infoLog *log.Logger
	errLog  *log.Logger
}

func NewGutHubLogger(out, outErr io.Writer, prefix string, flag int) *GutHubLogger {
	return &GutHubLogger{
		infoLog: log.New(out, prefix, flag),
		errLog:  log.New(outErr, prefix, flag),
	}
}

func (l *GutHubLogger) Info(v ...any) {
	l.infoLog.Print("INFO: ", fmt.Sprintln(v...))
}

func (l *GutHubLogger) Error(v ...any) {
	l.errLog.Print("ERROR: ", fmt.Sprintln(v...))
}

func (l *GutHubLogger) Debug(v ...any) {
	l.infoLog.Print(">>> DEBUG: ", fmt.Sprintln(v...))
}
