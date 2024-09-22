package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/minhnghia2k3/personal-blog/internal/handlers"
	"github.com/minhnghia2k3/personal-blog/internal/middlewares"
)

func RegisterArticleRoutes(r chi.Router, handler *handlers.ArticleHandler) {
	m := middlewares.New(handler.Service)

	r.Get("/", handler.GetAllArticles)

	r.Route("/articles", func(r chi.Router) {
		r.Get("/", handler.NewArticle)
		r.Post("/", handler.CreateArticle)
		r.Route("/{articleID}", func(r chi.Router) {
			r.Use(m.ArticleCtx)
			r.Get("/", handler.GetArticle)
			r.Put("/", handler.UpdateArticle)
			r.Delete("/", handler.DeleteArticle)
			r.Get("/edit", handler.EditArticle)
		})
	})
}
