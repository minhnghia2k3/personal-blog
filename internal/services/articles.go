package services

import (
	"github.com/minhnghia2k3/personal-blog/internal/models"
	"github.com/minhnghia2k3/personal-blog/internal/repositories"
	"time"
)

type ArticleService struct {
	repo repositories.ArticleRepository
}

func NewArticleService(repo repositories.ArticleRepository) *ArticleService {
	return &ArticleService{repo: repo}
}

func (s *ArticleService) GetArticleList() ([]*models.Article, error) {
	return s.repo.GetAll()

}

func (s *ArticleService) GetArticleById(id string) (*models.Article, error) {
	return s.repo.GetByID(id)
}

func (s *ArticleService) CreateArticle(article *models.Article) error {
	return s.repo.Create(article)
}

func (s *ArticleService) UpdateArticle(articleID string, article *models.Article) error {
	record, err := s.GetArticleById(articleID)
	if err != nil {
		return err
	}

	record.Title = article.Title
	record.Content = article.Content
	record.MinRead = article.MinRead
	record.UpdatedAt = time.Now()

	err = s.repo.Update(articleID, record)
	if err != nil {
		return err
	}
	return nil
}

func (s *ArticleService) DeleteArticle(id string) error {
	return s.repo.Delete(id)
}
