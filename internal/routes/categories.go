package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/minhnghia2k3/personal-blog/internal/handlers"
)

func RegisterCategoryRoutes(r *chi.Mux, h *handlers.CategoryHandlers) {
	r.Route("/categories", func(r chi.Router) {
		r.Get("/", h.NewCategory)
		r.Post("/", h.CreateCategory)
		r.Route("/{categoryID}", func(r chi.Router) {
			r.Get("/", h.EditCategory)
			r.Put("/", h.UpdateCategory)
			r.Delete("/", h.DeleteCategory)
		})
	})
}
