package repositories

import (
	"github.com/minhnghia2k3/personal-blog/internal/models"
)

type ArticleRepository interface {
	GetAll() ([]*models.Article, error)
	GetByID(articleID string) (*models.Article, error)
	Create(article *models.Article) error
	Update(articleID string, article *models.Article) error
	Delete(articleID string) error
}
