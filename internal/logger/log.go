package logger

import (
	"log/slog"
	"os"
)

type JSONLogger struct {
}

func New() *JSONLogger {
	return &JSONLogger{}
}

func (l *JSONLogger) DefaultLog() {
	appEnv := os.Getenv("APP_ENV")

	level := &slog.LevelVar{}

	opts := &slog.HandlerOptions{
		Level: level,
	}

	handler := slog.NewJSONHandler(os.Stdout, opts)

	switch appEnv {
	case "production":
		level.Set(slog.LevelInfo)
	case "development":
		level.Set(slog.LevelDebug)
	}

	logger := slog.New(handler)
	slog.SetDefault(logger)
}
