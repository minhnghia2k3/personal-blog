package dto

import "github.com/minhnghia2k3/personal-blog/internal/models"

type Pagination struct {
	Page   int    `json:"page"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
	Search string `json:"search"`
}

type Metadata struct {
	TotalCount  int `json:"total_count"`
	CurrentPage int `json:"current_page"`
	PageSize    int `json:"page_size"`
	TotalPages  int `json:"total_pages"`
}

type ArticleResponse struct {
	Article  []*models.Article `json:"article"`
	Metadata Metadata          `json:"metadata"`
}

type ArticleCategoriesResponse struct {
	ArticleCategories []*models.ArticleCategories
	Metadata          Metadata
}
