package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/minhnghia2k3/personal-blog/internal/handlers"
	"github.com/minhnghia2k3/personal-blog/internal/middlewares"
)

func RegisterArticleRoutes(r chi.Router, articleHandler *handlers.ArticleHandler, categoryHandler *handlers.CategoryHandlers) {
	m := middlewares.New(articleHandler.Service, categoryHandler.Service)

	r.Route("/articles", func(r chi.Router) {
		r.Use(m.CategoriesCtx)
		r.Get("/", articleHandler.NewArticle)
		r.Post("/", articleHandler.CreateArticle)
		r.Route("/{articleID}", func(r chi.Router) {
			r.Use(m.ArticleCtx)
			r.Get("/", articleHandler.GetArticle)
			r.Put("/", articleHandler.UpdateArticle)
			r.Delete("/", articleHandler.DeleteArticle)
			r.Get("/edit", articleHandler.EditArticle)
		})
	})
}
