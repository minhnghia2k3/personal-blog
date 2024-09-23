package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/minhnghia2k3/personal-blog/internal/helpers"
	"github.com/minhnghia2k3/personal-blog/internal/models"
	"github.com/minhnghia2k3/personal-blog/internal/services"
	"html/template"
	"net/http"
	"time"
)

type CategoryHandlers struct {
	Service *services.CategoryService
}

func NewCategoryHandlers(service *services.CategoryService) *CategoryHandlers {
	return &CategoryHandlers{service}
}

// CreateCategory handler that create new category.
func (h *CategoryHandlers) CreateCategory(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")

	category := &models.Category{
		Name:      name,
		CreatedAt: time.Now(),
	}

	err := h.Service.CreateCategory(category)
	helpers.HttpCatch(w, err)

	http.Redirect(w, r, "/", http.StatusFound)
}

// UpdateCategory handler that update a category.
func (h *CategoryHandlers) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	categoryID := chi.URLParam(r, "id")
	name := r.FormValue("name")

	category := &models.Category{
		Name: name,
	}
	err := h.Service.UpdateCategory(categoryID, category)
	helpers.HttpCatch(w, err)

	http.Redirect(w, r, "/", http.StatusFound)
}

// DeleteCategory handler that delete a category.
func (h *CategoryHandlers) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	categoryID := chi.URLParam(r, "id")
	err := h.Service.DeleteCategory(categoryID)
	helpers.HttpCatch(w, err)

	http.Redirect(w, r, "/", http.StatusOK)
}

// NewCategory will parse and render create category form.
func (h *CategoryHandlers) NewCategory(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./ui/html/base.html", "./ui/html/pages/create_category.html")
	helpers.HttpCatch(w, err)

	err = tmpl.Execute(w, nil)
	helpers.HttpCatch(w, err)
}

// EditCategory will parse and render edit category form with its value.
func (h *CategoryHandlers) EditCategory(w http.ResponseWriter, r *http.Request) {
	categoryID := chi.URLParam(r, "id")
	category, err := h.Service.GetCategoryByID(categoryID)
	helpers.HttpCatch(w, err)

	tmpl, err := template.ParseFiles("./ui/html/base.html", "./ui/html/pages/edit_category.html")
	helpers.HttpCatch(w, err)

	err = tmpl.Execute(w, category)
	helpers.HttpCatch(w, err)
}
