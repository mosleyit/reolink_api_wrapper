package reolink

import (
	"fmt"
	"io"
	"log"
	"os"
)

// Logger is the interface for logging in the Reolink client.
// Implement this interface to provide custom logging behavior.
type Logger interface {
	// Debug logs a debug message.
	Debug(msg string, args ...interface{})
	// Info logs an informational message.
	Info(msg string, args ...interface{})
	// Warn logs a warning message.
	Warn(msg string, args ...interface{})
	// Error logs an error message.
	Error(msg string, args ...interface{})
}

// NoOpLogger is a logger that does nothing.
// This is the default logger used by the client.
type NoOpLogger struct{}

// Debug does nothing.
func (l *NoOpLogger) Debug(msg string, args ...interface{}) {}

// Info does nothing.
func (l *NoOpLogger) Info(msg string, args ...interface{}) {}

// Warn does nothing.
func (l *NoOpLogger) Warn(msg string, args ...interface{}) {}

// Error does nothing.
func (l *NoOpLogger) Error(msg string, args ...interface{}) {}

// StdLogger is a simple logger that writes to an io.Writer.
// It uses the standard library's log package.
type StdLogger struct {
	debug *log.Logger
	info  *log.Logger
	warn  *log.Logger
	err   *log.Logger
}

// NewStdLogger creates a new StdLogger that writes to the given writer.
// If writer is nil, it defaults to os.Stderr.
func NewStdLogger(writer io.Writer) *StdLogger {
	if writer == nil {
		writer = os.Stderr
	}

	return &StdLogger{
		debug: log.New(writer, "[DEBUG] ", log.LstdFlags),
		info:  log.New(writer, "[INFO]  ", log.LstdFlags),
		warn:  log.New(writer, "[WARN]  ", log.LstdFlags),
		err:   log.New(writer, "[ERROR] ", log.LstdFlags),
	}
}

// Debug logs a debug message.
func (l *StdLogger) Debug(msg string, args ...interface{}) {
	if len(args) > 0 {
		l.debug.Printf(msg, args...)
	} else {
		l.debug.Print(msg)
	}
}

// Info logs an informational message.
func (l *StdLogger) Info(msg string, args ...interface{}) {
	if len(args) > 0 {
		l.info.Printf(msg, args...)
	} else {
		l.info.Print(msg)
	}
}

// Warn logs a warning message.
func (l *StdLogger) Warn(msg string, args ...interface{}) {
	if len(args) > 0 {
		l.warn.Printf(msg, args...)
	} else {
		l.warn.Print(msg)
	}
}

// Error logs an error message.
func (l *StdLogger) Error(msg string, args ...interface{}) {
	if len(args) > 0 {
		l.err.Printf(msg, args...)
	} else {
		l.err.Print(msg)
	}
}

// LevelLogger is a logger that supports log levels.
// Messages below the configured level are not logged.
type LevelLogger struct {
	level  LogLevel
	logger Logger
}

// LogLevel represents the severity level of a log message.
type LogLevel int

const (
	// LogLevelDebug is the debug log level.
	LogLevelDebug LogLevel = iota
	// LogLevelInfo is the info log level.
	LogLevelInfo
	// LogLevelWarn is the warn log level.
	LogLevelWarn
	// LogLevelError is the error log level.
	LogLevelError
	// LogLevelNone disables all logging.
	LogLevelNone
)

// String returns the string representation of the log level.
func (l LogLevel) String() string {
	switch l {
	case LogLevelDebug:
		return "DEBUG"
	case LogLevelInfo:
		return "INFO"
	case LogLevelWarn:
		return "WARN"
	case LogLevelError:
		return "ERROR"
	case LogLevelNone:
		return "NONE"
	default:
		return fmt.Sprintf("UNKNOWN(%d)", l)
	}
}

// NewLevelLogger creates a new LevelLogger with the given level and underlying logger.
func NewLevelLogger(level LogLevel, logger Logger) *LevelLogger {
	return &LevelLogger{
		level:  level,
		logger: logger,
	}
}

// Debug logs a debug message if the level is Debug or higher.
func (l *LevelLogger) Debug(msg string, args ...interface{}) {
	if l.level <= LogLevelDebug {
		l.logger.Debug(msg, args...)
	}
}

// Info logs an informational message if the level is Info or higher.
func (l *LevelLogger) Info(msg string, args ...interface{}) {
	if l.level <= LogLevelInfo {
		l.logger.Info(msg, args...)
	}
}

// Warn logs a warning message if the level is Warn or higher.
func (l *LevelLogger) Warn(msg string, args ...interface{}) {
	if l.level <= LogLevelWarn {
		l.logger.Warn(msg, args...)
	}
}

// Error logs an error message if the level is Error or higher.
func (l *LevelLogger) Error(msg string, args ...interface{}) {
	if l.level <= LogLevelError {
		l.logger.Error(msg, args...)
	}
}
