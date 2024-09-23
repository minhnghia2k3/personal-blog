package repositories

import (
	"github.com/minhnghia2k3/personal-blog/internal/models"
)

type ArticleRepository interface {
	GetAll() ([]*models.Article, error)
	GetByID(articleID string) (*models.Article, error)
	Create(article *models.Article) (*models.Article, error)
	Update(articleID string, article *models.Article) error
	Delete(articleID string) error

	AddCategory(articleID, categoryID string) error
	RemoveCategory(articleID, categoryID string) error
	GetCategoriesByArticle(articleID string) ([]*models.Category, error)
}

type CategoryRepository interface {
	GetAll() ([]*models.Category, error)
	GetByID(categoryID string) (*models.Category, error)
	GetByName(name string) (*models.Category, error)
	Create(category *models.Category) error
	Update(categoryID string, category *models.Category) error
	Delete(categoryID string) error
}
