package main

import (
	"github.com/minhnghia2k3/personal-blog/internal"
	"github.com/minhnghia2k3/personal-blog/internal/config"
	"github.com/minhnghia2k3/personal-blog/internal/database"
	"github.com/minhnghia2k3/personal-blog/internal/logger"
	"github.com/minhnghia2k3/personal-blog/internal/routes"
	"log/slog"
)

func main() {
	// Load config
	cfg := config.Load()

	// Initialize logger
	consoleLog := logger.ConsoleLogger{}
	err := consoleLog.NewLogger(slog.LevelInfo)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	defer consoleLog.CloseLogger()

	// Initialize application
	app := internal.NewApplication(cfg)

	// Initialize routes
	r := routes.Routes()

	// Connect database
	db, err := database.ConnectDB(cfg)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	defer db.Close()

	// Serve HTTP Server
	err = app.Serve(r)
	if err != nil {
		slog.Error(err.Error())
	}
}
