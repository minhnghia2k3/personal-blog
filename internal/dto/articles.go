package dto

import "github.com/minhnghia2k3/personal-blog/internal/models"

type CreateArticle struct {
	Article        *models.Article
	CategoriesName []string
}

type EditArticle struct {
	Article        *models.Article
	CategoriesName []string
	ListCategories []*models.Category
}
