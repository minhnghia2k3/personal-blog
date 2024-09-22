package handlers

import (
	"fmt"
	"github.com/minhnghia2k3/personal-blog/internal/helpers"
	"github.com/minhnghia2k3/personal-blog/internal/models"
	"github.com/minhnghia2k3/personal-blog/internal/services"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

type ArticleHandler struct {
	Service *services.ArticleService
}

func NewArticleHandler(service *services.ArticleService) *ArticleHandler {
	return &ArticleHandler{Service: service}
}

func (h *ArticleHandler) GetAllArticles(w http.ResponseWriter, r *http.Request) {
	records, err := h.Service.GetArticleList()
	helpers.Catch(err)

	t, _ := template.ParseFiles("ui/html/base.html", "ui/html/pages/index.html")
	err = t.Execute(w, records)
	helpers.Catch(err)
}

func (h *ArticleHandler) NewArticle(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("ui/html/base.html", "ui/html/pages/create_article.html")
	err := t.Execute(w, nil)
	helpers.Catch(err)
}

func (h *ArticleHandler) CreateArticle(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	content := r.FormValue("content")
	minRead, err := strconv.Atoi(r.FormValue("min_read"))
	helpers.Catch(err)

	article := &models.Article{
		Title:     title,
		Content:   template.HTML(content),
		MinRead:   minRead,
		CreatedAt: time.Now(),
	}

	err = h.Service.CreateArticle(article)
	helpers.Catch(err)
	http.Redirect(w, r, "/", http.StatusFound)
}

func (h *ArticleHandler) GetArticle(w http.ResponseWriter, r *http.Request) {
	article := r.Context().Value("article").(*models.Article)
	t, _ := template.ParseFiles("ui/html/base.html", "ui/html/pages/article.html")
	err := t.Execute(w, article)
	helpers.Catch(err)
}

func (h *ArticleHandler) EditArticle(w http.ResponseWriter, r *http.Request) {
	article := r.Context().Value("article").(*models.Article)
	t, _ := template.ParseFiles("ui/html/base.html", "ui/html/pages/edit.html")
	err := t.Execute(w, article)
	helpers.Catch(err)
}

func (h *ArticleHandler) UpdateArticle(w http.ResponseWriter, r *http.Request) {
	article := r.Context().Value("article").(*models.Article)

	title := r.FormValue("title")
	content := r.FormValue("content")
	minRead, err := strconv.Atoi(r.FormValue("min_read"))
	helpers.Catch(err)

	newArticle := &models.Article{
		Title:     title,
		Content:   template.HTML(content),
		MinRead:   minRead,
		UpdatedAt: time.Now(),
	}
	err = h.Service.UpdateArticle(strconv.Itoa(article.ID), newArticle)
	helpers.Catch(err)

	http.Redirect(w, r, fmt.Sprintf("/articles/%d", article.ID), http.StatusFound)
}

func (h *ArticleHandler) DeleteArticle(w http.ResponseWriter, r *http.Request) {
	article := r.Context().Value("article").(*models.Article)
	err := h.Service.DeleteArticle(strconv.Itoa(article.ID))
	helpers.Catch(err)

	http.Redirect(w, r, "/", http.StatusOK)
}
