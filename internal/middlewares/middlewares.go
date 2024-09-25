package middlewares

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/minhnghia2k3/personal-blog/internal/helpers"
	"github.com/minhnghia2k3/personal-blog/internal/services"
	"net/http"
)

type constant string

const (
	ArticleConstant    constant = "article"
	CategoriesConstant constant = "categories"
)

type Middleware struct {
	articleService  *services.ArticleService
	categoryService *services.CategoryService
}

func New(articleService *services.ArticleService, categoryService *services.CategoryService) *Middleware {
	return &Middleware{articleService, categoryService}
}

// =========================== Middlewares ===========================

func ChangeMethod(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			switch method := r.PostFormValue("_method"); method {
			case http.MethodPut:
				fallthrough
			case http.MethodPatch:
				fallthrough
			case http.MethodDelete:
				r.Method = method
			default:
			}
		}
		next.ServeHTTP(w, r)
	})
}

func (m *Middleware) ArticleCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		articleID := chi.URLParam(r, "articleID")

		// Fetch article from database
		article, err := m.articleService.GetArticleById(articleID)
		helpers.HttpCatch(w, http.StatusInternalServerError, err)

		ctx := context.WithValue(r.Context(), ArticleConstant, article)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (m *Middleware) CategoriesCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		categories, err := m.categoryService.ListCategories()
		helpers.HttpCatch(w, http.StatusInternalServerError, err)

		ctx := context.WithValue(r.Context(), CategoriesConstant, categories)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
