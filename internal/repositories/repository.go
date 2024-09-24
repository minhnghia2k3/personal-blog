package repositories

import (
	"github.com/minhnghia2k3/personal-blog/internal/dto"
	"github.com/minhnghia2k3/personal-blog/internal/models"
)

type ArticleRepository interface {
	GetAll(pagination dto.Pagination) (*dto.ArticleResponse, error)
	GetByID(string) (*models.Article, error)
	Create(*models.Article) (*models.Article, error)
	Update(string, *models.Article) error
	Delete(string) error

	AddCategory(string, string) error
	RemoveCategory(string, string) error
	GetCategoriesByArticle(string) ([]*models.Category, error)
}

type CategoryRepository interface {
	GetAll() ([]*models.Category, error)
	GetByID(string) (*models.Category, error)
	GetByName(string) (*models.Category, error)
	Create(*models.Category) error
	Update(string, *models.Category) error
	Delete(string) error
}
