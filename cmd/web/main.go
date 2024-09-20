package main

import (
	"github.com/minhnghia2k3/personal-blog/internal"
	"github.com/minhnghia2k3/personal-blog/internal/config"
	logger "github.com/minhnghia2k3/personal-blog/internal/logger"
	"github.com/minhnghia2k3/personal-blog/internal/routes"
	"log/slog"
)

func main() {
	// Load config
	cfg := config.Load()

	// Initialize logger
	err := logger.NewLogger(slog.LevelInfo)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	defer logger.CloseLogger()

	// Initialize application
	app := internal.NewApplication(cfg)

	// Initialize routes
	r := routes.Routes()
	// Serve HTTP Server
	err = app.Serve(r)
	if err != nil {
		slog.Error(err.Error())
	}
}
