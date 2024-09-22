package repositories

import (
	"database/sql"
	"github.com/minhnghia2k3/personal-blog/internal/models"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) GetAll() ([]*models.Article, error) {
	var articles []*models.Article
	stmt := `SELECT id, title, content, min_read, created_at, updated_at
FROM articles`
	query, err := r.DB.Prepare(stmt)
	if err != nil {
		return nil, err
	}
	defer query.Close()

	rows, err := query.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var article models.Article
		err = rows.Scan(&article.ID, &article.Title, &article.Content, &article.MinRead, &article.CreatedAt, &article.UpdatedAt)
		if err != nil {
			return nil, err
		}

		articles = append(articles, &article)
	}

	return articles, nil
}

func (r *Repository) GetByID(articleID string) (*models.Article, error) {
	var article models.Article
	stmt := `SELECT id, title, content, min_read, created_at, updated_at
FROM articles WHERE id = $1`

	query, err := r.DB.Prepare(stmt)
	defer query.Close()

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

func (r *Repository) Create(article *models.Article) error {
	stmt := `INSERT INTO articles(title, content, min_read, created_at)
VALUES($1, $2, $3, $4)`
	query, err := r.DB.Prepare(stmt)
	if err != nil {
		return err
	}

	_, err = query.Exec(article.Title, article.Content, article.MinRead, article.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Update(articleID string, article *models.Article) error {
	stmt := `UPDATE articles SET title = $1, content = $2, min_read = $3, updated_at = $4
WHERE id = $5`
	query, err := r.DB.Prepare(stmt)

	if err != nil {
		return err
	}
	defer query.Close()

	_, err = query.Exec(article.Title, article.Content, article.MinRead, article.UpdatedAt, articleID)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Delete(articleID string) error {
	stmt := `DELETE FROM articles WHERE id = $1`
	query, err := r.DB.Prepare(stmt)

	if err != nil {
		return err
	}
	defer query.Close()

	_, err = query.Exec(articleID)
	if err != nil {
		return err
	}

	return nil
}
