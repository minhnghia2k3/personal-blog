package repositories

import (
	"database/sql"
	"github.com/minhnghia2k3/personal-blog/internal/dto"
	"github.com/minhnghia2k3/personal-blog/internal/models"
	"math"
)

type PostgresArticleRepository struct {
	DB *sql.DB
}

func NewPostgresArticleRepository(db *sql.DB) *PostgresArticleRepository {
	return &PostgresArticleRepository{DB: db}
}

func handleError(query *sql.Stmt) {
	if query != nil {
		_ = query.Close()
	}
}

func (r *PostgresArticleRepository) GetAll(p dto.Pagination) (*dto.ArticleResponse, error) {
	var articles []*models.Article
	var totalCount int

	stmt := `SELECT id, title, content, min_read, created_at, updated_at FROM articles`

	countStmt := `SELECT COUNT(*) FROM articles`

	var params []interface{}
	params = append(params, p.Limit, p.Offset)

	if p.Search != "" {
		p.Search = "%" + p.Search + "%"
		stmt += ` WHERE(title ILIKE $3 OR content ILIKE $3)`
		params = append(params, p.Search)
	}

	stmt += ` ORDER BY created_at desc LIMIT $1 OFFSET $2`

	// Get total records
	err := r.DB.QueryRow(countStmt).Scan(&totalCount)
	if err != nil {
		return nil, err
	}

	// Get the paginated articles
	query, err := r.DB.Prepare(stmt)
	if err != nil {
		return nil, err
	}
	defer handleError(query)

	rows, err := query.Query(params...)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	// Process article rows
	for rows.Next() {
		var article models.Article
		err = rows.Scan(&article.ID, &article.Title, &article.Content, &article.MinRead, &article.CreatedAt, &article.UpdatedAt)
		if err != nil {
			return nil, err
		}

		articles = append(articles, &article)
	}

	// Calculate total pages
	totalPages := int(math.Ceil(float64(totalCount) / float64(p.Limit)))

	// Prepare response with metadata
	response := &dto.ArticleResponse{
		Article: articles,
		Metadata: dto.Metadata{
			TotalCount:  totalCount,
			CurrentPage: p.Page,
			PageSize:    p.Limit,
			TotalPages:  totalPages,
		},
	}

	return response, nil
}

func (r *PostgresArticleRepository) GetByID(articleID string) (*models.Article, error) {
	var article models.Article
	stmt := `SELECT id, title, content, min_read, created_at, updated_at
FROM articles WHERE id = $1`

	query, err := r.DB.Prepare(stmt)
	defer handleError(query)

	if err != nil {
		return nil, err
	}

	row := query.QueryRow(articleID)
	err = row.Scan(&article.ID, &article.Title, &article.Content, &article.MinRead, &article.CreatedAt, &article.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &article, nil
}

func (r *PostgresArticleRepository) Create(article *models.Article) (*models.Article, error) {
	result := new(models.Article)

	stmt := `INSERT INTO articles(title, content, min_read, created_at) VALUES($1, $2, $3, $4)
RETURNING id, title, content, min_read, created_at`
	query, err := r.DB.Prepare(stmt)
	if err != nil {
		return nil, err
	}
	defer handleError(query)

	// Execute the query and scan the returned values into the article
	err = query.QueryRow(article.Title, article.Content, article.MinRead, article.CreatedAt).Scan(
		&result.ID,
		&result.Title,
		&result.Content,
		&result.MinRead,
		&result.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *PostgresArticleRepository) Update(articleID string, article *models.Article) error {
	stmt := `UPDATE articles SET title = $1, content = $2, min_read = $3, updated_at = $4
WHERE id = $5`
	query, err := r.DB.Prepare(stmt)

	if err != nil {
		return err
	}
	defer handleError(query)

	_, err = query.Exec(article.Title, article.Content, article.MinRead, article.UpdatedAt, articleID)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresArticleRepository) Delete(articleID string) error {
	stmt := `DELETE FROM articles WHERE id = $1`
	query, err := r.DB.Prepare(stmt)

	if err != nil {
		return err
	}
	defer handleError(query)

	_, err = query.Exec(articleID)
	if err != nil {
		return err
	}

	return nil
}

// AddCategory associates category with article.
func (r *PostgresArticleRepository) AddCategory(articleID, categoryID string) error {
	stmt := `INSERT INTO article_categories(article_id, category_id) VALUES ($1, $2)`

	query, err := r.DB.Prepare(stmt)
	if err != nil {
		return err
	}
	defer handleError(query)

	_, err = query.Exec(articleID, categoryID)
	if err != nil {
		return err
	}
	return nil
}

// RemoveCategory removes category from article.
func (r *PostgresArticleRepository) RemoveCategory(articleID, categoryID string) error {
	stmt := `DELETE FROM article_categories WHERE article_id = $1 AND category_id = $2`

	query, err := r.DB.Prepare(stmt)
	if err != nil {
		return err
	}
	defer handleError(query)

	_, err = query.Exec(articleID, categoryID)
	if err != nil {
		return err
	}

	return nil
}

// GetCategoriesByArticle fetches all categories that associated with the article.
func (r *PostgresArticleRepository) GetCategoriesByArticle(articleID string) ([]*models.Category, error) {
	var categories []*models.Category
	stmt := `SELECT id, name FROM categories c
INNER JOIN article_categories ac ON c.ID = ac.category_id
WHERE ac.article_id = $1`

	query, err := r.DB.Prepare(stmt)
	if err != nil {
		return nil, err
	}
	defer handleError(query)

	rows, err := query.Query(articleID)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	for rows.Next() {
		var category models.Category
		err = rows.Scan(&category.ID, &category.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}
	return categories, nil
}
