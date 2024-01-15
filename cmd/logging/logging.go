package logging

import (
	"fmt"
	"io"
	"log"
)

// Logger is an interface for logging.
type Logger interface {
	Info(v ...any)
	Error(v ...any)
	Debug(v ...any)
}

// GutHubLogger is a logger for GutHub.
type GutHubLogger struct {
	infoLog  *log.Logger
	errLog   *log.Logger
	debugLog *log.Logger
}

// NewGutHubLogger returns a new GutHubLogger.
func NewGutHubLogger(infoOut, debugOut, errOut io.Writer, prefix string, flag int) *GutHubLogger {
	return &GutHubLogger{
		infoLog:  log.New(infoOut, prefix+"INFO: ", flag),
		debugLog: log.New(debugOut, prefix+">>> DEBUG: ", flag),
		errLog:   log.New(errOut, prefix+"ERROR: ", flag),
	}
}

// Info logs info messages to stdout.
func (l *GutHubLogger) Info(v ...any) {
	l.infoLog.Print(fmt.Sprintln(v...))
}

// Error logs error messages to stderr.
func (l *GutHubLogger) Error(v ...any) {
	l.errLog.Print(fmt.Sprintln(v...))
}

// Debug logs debug messages, with a special >>> prefix to make them stand out.
func (l *GutHubLogger) Debug(v ...any) {
	l.debugLog.Print(fmt.Sprintln(v...))
}
