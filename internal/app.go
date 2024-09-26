package internal

import (
	"fmt"
	"github.com/minhnghia2k3/personal-blog/internal/config"
	"log"
	"net/http"
	"time"
)

type Application struct {
	Config *config.Config
}

func NewApplication(cfg *config.Config) *Application {
	return &Application{
		Config: cfg,
	}
}

func (app *Application) Serve(handler http.Handler) error {
	srv := http.Server{
		Addr:         fmt.Sprintf(":%d", app.Config.Port),
		Handler:      handler,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Printf("Listening on port :%d\n", app.Config.Port)
	return srv.ListenAndServe()
}
