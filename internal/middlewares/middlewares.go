package middlewares

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/minhnghia2k3/personal-blog/internal/services"
	"net/http"
)

type constant string

const ArticleConstant constant = "article"

type Middleware struct {
	articleService *services.ArticleService
}

func New(articleService *services.ArticleService) *Middleware {
	return &Middleware{articleService}
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
		if err != nil {
			fmt.Println(err)
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		ctx := context.WithValue(r.Context(), ArticleConstant, article)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
