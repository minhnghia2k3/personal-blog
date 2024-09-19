package handlers

import (
	"fmt"
	"github.com/minhnghia2k3/personal-blog/cmd/config"
	"net/http"
)

func CheckHealthHandler(w http.ResponseWriter, r *http.Request) {
	cfg := config.Load()
	w.WriteHeader(http.StatusOK)

	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "stage: %s\n", cfg.Env)
	fmt.Fprintf(w, "version: %s\n", config.Version)
}
