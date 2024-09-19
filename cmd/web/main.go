package main

import (
	"github.com/minhnghia2k3/personal-blog/internal/config"
	"github.com/minhnghia2k3/personal-blog/internal/routes"
)

func main() {
	// Load config
	cfg := config.Load()

	// Define routes
	mux := routes.Routes()

	// Serve HTTP server
	Serve(cfg, mux)
}
