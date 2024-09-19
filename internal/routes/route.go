package routes

import (
	"github.com/minhnghia2k3/personal-blog/internal/handlers"
	"net/http"
)

func Routes() http.Handler {
	mux := http.NewServeMux()

	// healthcheck route
	mux.HandleFunc("GET /healthcheck", handlers.CheckHealthHandler)
	return mux
}
