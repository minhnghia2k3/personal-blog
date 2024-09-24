package handlers

import (
	"fmt"
	"github.com/minhnghia2k3/personal-blog/internal/dto"
	"github.com/minhnghia2k3/personal-blog/internal/helpers"
	"github.com/minhnghia2k3/personal-blog/internal/middlewares"
	"github.com/minhnghia2k3/personal-blog/internal/models"
	"github.com/minhnghia2k3/personal-blog/internal/services"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

type ArticleHandler struct {
	Service         *services.ArticleService
	CategoryService *services.CategoryService
}

func NewArticleHandler(service *services.ArticleService, categoryService *services.CategoryService) *ArticleHandler {
	return &ArticleHandler{
		Service:         service,
		CategoryService: categoryService}
}

// GetAllArticles will fetch all articles in database and render to the template.
func (h *ArticleHandler) GetAllArticles(w http.ResponseWriter, r *http.Request) {
	// Get query parameters
	p := helpers.GetPaginationValues(r)

	response, err := h.Service.GetArticleList(p)
	helpers.HttpCatch(w, err)

	// Execute data to template
	t, _ := template.ParseFiles("ui/html/base.html", "ui/html/pages/index.html")
	err = t.Execute(w, response)
	helpers.HttpCatch(w, err)
}

// GetArticle get article information and render to the template.
func (h *ArticleHandler) GetArticle(w http.ResponseWriter, r *http.Request) {
	article := r.Context().Value(middlewares.ArticleConstant).(*models.Article)
	categories, err := h.Service.GetCategoryList(strconv.Itoa(article.ID))
	helpers.HttpCatch(w, err)

	articleCategories := &models.ArticleCategories{
		Article:    article,
		Categories: categories,
	}

	t, _ := template.ParseFiles("ui/html/base.html", "ui/html/pages/article.html")
	err = t.Execute(w, articleCategories)
	helpers.HttpCatch(w, err)
}

// NewArticle will render create article form with categories data for multiple select.
func (h *ArticleHandler) NewArticle(w http.ResponseWriter, r *http.Request) {
	categories := r.Context().Value(middlewares.CategoriesConstant).([]*models.Category)
	t, _ := template.ParseFiles("ui/html/base.html", "ui/html/pages/create_article.html")
	err := t.Execute(w, categories)
	helpers.HttpCatch(w, err)
}

// CreateArticle is an action for creating new article. Will redirect to homepage If successfully.
func (h *ArticleHandler) CreateArticle(w http.ResponseWriter, r *http.Request) {
	// Parse from value
	title := r.FormValue("title")
	content := r.FormValue("content")
	minRead, err := strconv.Atoi(r.FormValue("min_read"))
	helpers.HttpCatch(w, err)
	categoryNames := r.Form["categories"]

	// Create new article
	article := &dto.CreateArticle{
		Article: &models.Article{
			Title:     title,
			Content:   template.HTML(content),
			MinRead:   minRead,
			CreatedAt: time.Now(),
		},
		CategoriesName: categoryNames,
	}
	err = h.Service.CreateArticle(article)
	helpers.HttpCatch(w, err)

	http.Redirect(w, r, "/", http.StatusFound)
}

// EditArticle will render edit article form.
func (h *ArticleHandler) EditArticle(w http.ResponseWriter, r *http.Request) {
	listCategories := r.Context().Value(middlewares.CategoriesConstant).([]*models.Category)
	article := r.Context().Value(middlewares.ArticleConstant).(*models.Article)

	articleCategories := &dto.EditArticle{
		Article:        article,
		ListCategories: listCategories,
	}

	t, _ := template.ParseFiles("ui/html/base.html", "ui/html/pages/edit.html")
	err := t.Execute(w, articleCategories)
	helpers.HttpCatch(w, err)
}

// UpdateArticle is a function for updating an article.
func (h *ArticleHandler) UpdateArticle(w http.ResponseWriter, r *http.Request) {
	article := r.Context().Value(middlewares.ArticleConstant).(*models.Article)

	title := r.FormValue("title")
	content := r.FormValue("content")
	minRead, err := strconv.Atoi(r.FormValue("min_read"))
	helpers.HttpCatch(w, err)
	categoryNames := r.Form["categories"]

	// Create new article_categories
	newArticle := &dto.EditArticle{
		Article: &models.Article{
			Title:     title,
			Content:   template.HTML(content),
			MinRead:   minRead,
			UpdatedAt: time.Now(),
		},
		CategoriesName: categoryNames,
	}
	err = h.Service.UpdateArticle(strconv.Itoa(article.ID), newArticle)
	helpers.HttpCatch(w, err)

	http.Redirect(w, r, fmt.Sprintf("/articles/%d", article.ID), http.StatusFound)
}

// DeleteArticle will get article and delete it. Will redirect to homepage If successfully.
func (h *ArticleHandler) DeleteArticle(w http.ResponseWriter, r *http.Request) {
	article := r.Context().Value(middlewares.ArticleConstant).(*models.Article)
	err := h.Service.DeleteArticle(strconv.Itoa(article.ID))
	helpers.HttpCatch(w, err)

	http.Redirect(w, r, "/", http.StatusOK)
}
