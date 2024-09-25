package dto

import "github.com/minhnghia2k3/personal-blog/internal/models"

type CreateArticle struct {
	Article        *models.Article
	CategoriesName []string `json:"categories" validate:"required"`
}

type EditArticle struct {
	Article        *models.Article
	CategoriesName []string `json:"categories" validate:"required"`
	ListCategories []*models.Category
}
