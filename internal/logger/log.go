package logger

import (
	"log/slog"
	"os"
)

type Logger interface {
	NewLogger(level slog.Level) error
	CloseLogger() error
}

// ConsoleLogger logs to standard output
type ConsoleLogger struct{}

// FileLogger logs to specified log file
type FileLogger struct {
	logFile *os.File
}

// NewLogger initialize console logger that log to standard output
func (l *ConsoleLogger) NewLogger(level slog.Level) error {
	logLevel := new(slog.LevelVar)
	logger := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel})
	slog.SetDefault(slog.New(logger))
	logLevel.Set(level)
	return nil
}

// NewLogger initialize file logger that log to specific log file
func (l *FileLogger) NewLogger(level slog.Level) error {
	var err error
	l.logFile, err = os.OpenFile(os.Getenv("LOG_FILE"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	var logLevel = new(slog.LevelVar)
	logger := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel})
	slog.SetDefault(slog.New(logger))

	logLevel.Set(level)
	return nil
}

// CloseLogger closes the log file if a file logger is used
func (l *FileLogger) CloseLogger() error {
	if l.logFile != nil {
		return l.logFile.Close()
	}
	return nil
}

// CloseLogger does not close anything
func (l *ConsoleLogger) CloseLogger() error {
	return nil
}
