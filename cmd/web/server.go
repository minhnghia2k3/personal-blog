package main

import (
	"fmt"
	"github.com/minhnghia2k3/personal-blog/internal/config"
	"log"
	"net/http"
	"time"
)

func Serve(cfg *config.Config, handler http.Handler) {
	srv := http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      handler,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	fmt.Println("Listening on port :8080")
	log.Fatal(srv.ListenAndServe())
}
