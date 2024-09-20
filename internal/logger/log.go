package logger

import (
	"log/slog"
	"os"
)

var logFile *os.File

func NewLogger(level slog.Level) error {
	logFile, err := os.OpenFile("./log/log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	var logLevel = new(slog.LevelVar)
	logger := slog.NewJSONHandler(logFile, &slog.HandlerOptions{Level: logLevel})
	slog.SetDefault(slog.New(logger))

	logLevel.Set(level)
	slog.Info("This is a slog")
	return nil
}

func CloseLogger() {
	if logFile != nil {
		logFile.Close()
	}
}
