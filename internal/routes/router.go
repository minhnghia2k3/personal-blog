package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/minhnghia2k3/personal-blog/internal/handlers"
	"github.com/minhnghia2k3/personal-blog/internal/middlewares"
	"github.com/minhnghia2k3/personal-blog/ui"
	"net/http"
)

func Routes(articleHandler *handlers.ArticleHandler, categoryHandler *handlers.CategoryHandlers, imageHandler *handlers.ImageHandler) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middlewares.ChangeMethod)
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServerFS(ui.StaticFS)))

	r.Get("/", articleHandler.GetAllArticles)
	RegisterArticleRoutes(r, articleHandler, categoryHandler)
	RegisterCategoryRoutes(r, categoryHandler)
	RegisterImageRoutes(r, imageHandler)
	return r
}
