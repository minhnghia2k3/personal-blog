package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/minhnghia2k3/personal-blog/internal/handlers"
	"github.com/minhnghia2k3/personal-blog/internal/middlewares"
)

func Routes(articleHandler *handlers.ArticleHandler, imageHandler *handlers.ImageHandler) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middlewares.ChangeMethod)
	RegisterArticleRoutes(r, articleHandler)
	RegisterImageRoutes(r, imageHandler)
	return r
}
