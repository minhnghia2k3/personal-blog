package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/minhnghia2k3/personal-blog/internal/handlers"
)

func RegisterImageRoutes(r chi.Router, handler *handlers.ImageHandler) {
	r.Post("/upload", handler.UploadHandler)
	r.Get("/images/*", handler.ServeImages)
}
