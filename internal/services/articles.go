package services

import (
	"errors"
	"fmt"
	"github.com/minhnghia2k3/personal-blog/internal/dto"
	"github.com/minhnghia2k3/personal-blog/internal/helpers"
	"github.com/minhnghia2k3/personal-blog/internal/models"
	"github.com/minhnghia2k3/personal-blog/internal/repositories"
	"log"
	"strconv"
	"time"
)

type ArticleService struct {
	repo         repositories.ArticleRepository
	categoryRepo repositories.CategoryRepository
}

func NewArticleService(repo repositories.ArticleRepository, categoryRepo repositories.CategoryRepository) *ArticleService {
	return &ArticleService{repo: repo, categoryRepo: categoryRepo}
}

// GetArticleList will fetch all articles from the database.
func (s *ArticleService) GetArticleList(p dto.Pagination) (*dto.ArticleCategoriesResponse, error) {
	var articleCategories []*models.ArticleCategories

	// Calculate offset
	offset := (p.Page - 1) * p.Limit
	p.Offset = offset

	// Get article list
	articles, err := s.repo.GetAll(p)
	if err != nil {
		log.Printf("Failed to get articles from database: %v", err)
		return nil, errors.New("failed to get articles")
	}

	// Get article categories
	for _, article := range articles.Article {
		categories, err := s.repo.GetCategoriesByArticle(strconv.Itoa(article.ID))
		if err != nil {
			log.Printf("Failed to get categories from database: %v", err)
			return nil, fmt.Errorf("failed to get categories by article")
		}

		articleCategories = append(articleCategories, &models.ArticleCategories{
			Article:    article,
			Categories: categories,
		})
	}

	// Prepare response data
	response := &dto.ArticleCategoriesResponse{
		ArticleCategories: articleCategories,
		Metadata:          articles.Metadata,
	}

	return response, nil
}

// GetArticleById will fetch article by ID.
func (s *ArticleService) GetArticleById(id string) (*models.Article, error) {
	return s.repo.GetByID(id)
}

// CreateArticle will create new article record along with its categories.
func (s *ArticleService) CreateArticle(article *dto.CreateArticle) error {
	// Create new article first
	result, err := s.repo.Create(article.Article)
	if err != nil {
		return err
	}

	// Assign each category to the created article
	for _, name := range article.CategoriesName {
		// Find category by name
		c, err := s.categoryRepo.GetByName(name)
		if err != nil {
			return err
		}
		err = s.repo.AddCategory(strconv.Itoa(result.ID), strconv.Itoa(c.ID))
		if err != nil {
			return err
		}
	}

	return nil
}

// UpdateArticle will update an existing record.
func (s *ArticleService) UpdateArticle(articleID string, articleDto *dto.EditArticle) error {
	result, err := s.GetArticleById(articleID)
	if err != nil {
		return fmt.Errorf("failed to get article: %w", err)
	}

	result.Title = articleDto.Article.Title
	result.Content = articleDto.Article.Content
	result.MinRead = articleDto.Article.MinRead
	result.UpdatedAt = time.Now()

	// Get categories for duplicate checking.
	categories, err := s.GetCategoryList(strconv.Itoa(result.ID))
	if err != nil {
		return fmt.Errorf("failed to get categories: %w", err)
	}

	// Assign each category to the created article
	for _, name := range articleDto.CategoriesName {
		c, err := s.categoryRepo.GetByName(name)
		if err != nil {
			return fmt.Errorf("failed to get category by name '%s': %w", name, err)
		}

		// Check if the category is already assigned
		if !helpers.ContainsCategory(categories, c) {
			if err := s.repo.AddCategory(strconv.Itoa(result.ID), strconv.Itoa(c.ID)); err != nil {
				return fmt.Errorf("failed to add category '%s' to article: %w", name, err)
			}
		}
	}

	if err := s.repo.Update(articleID, result); err != nil {
		return fmt.Errorf("failed to update article: %w", err)
	}
	return nil
}

// DeleteArticle will delete an article.
func (s *ArticleService) DeleteArticle(id string) error {
	return s.repo.Delete(id)
}

// AddCategory associates article with category.
func (s *ArticleService) AddCategory(articleID, categoryID string) error {
	return s.repo.AddCategory(articleID, categoryID)
}

// RemoveCategory removes category from article.
func (s *ArticleService) RemoveCategory(articleID, categoryID string) error {
	return s.repo.RemoveCategory(articleID, categoryID)
}

// GetCategoryList fetch all category from specified article.
func (s *ArticleService) GetCategoryList(articleID string) ([]*models.Category, error) {
	return s.repo.GetCategoriesByArticle(articleID)
}
