package services

import (
	"github.com/minhnghia2k3/personal-blog/internal/models"
	"github.com/minhnghia2k3/personal-blog/internal/repositories"
)

type CategoryService struct {
	repo *repositories.PostgresCategoryRepository
}

func NewCategoryService(repo *repositories.PostgresCategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

// ListCategories fetches all categories in the database.
func (s *CategoryService) ListCategories() ([]*models.Category, error) {
	return s.repo.GetAll()
}

// GetCategoryByID fetches category by ID.
func (s *CategoryService) GetCategoryByID(id string) (*models.Category, error) {
	return s.repo.GetByID(id)
}

func (s *CategoryService) GetCategoryByName(name string) (*models.Category, error) {
	return s.repo.GetByName(name)
}

// CreateCategory creates new category.
func (s *CategoryService) CreateCategory(category *models.Category) error {
	return s.repo.Create(category)
}

// UpdateCategory updates a category.
func (s *CategoryService) UpdateCategory(categoryID string, category *models.Category) error {
	record, err := s.GetCategoryByID(categoryID)
	if err != nil {
		return err
	}

	record.Name = category.Name

	return s.repo.Update(categoryID, record)
}

// DeleteCategory delete a category.
func (s *CategoryService) DeleteCategory(categoryID string) error {
	return s.repo.Delete(categoryID)
}
