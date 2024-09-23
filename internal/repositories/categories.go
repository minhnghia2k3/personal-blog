package repositories

import (
	"database/sql"
	"github.com/minhnghia2k3/personal-blog/internal/models"
)

type PostgresCategoryRepository struct {
	DB *sql.DB
}

func NewPostgresCategoryRepository(db *sql.DB) *PostgresCategoryRepository {
	return &PostgresCategoryRepository{DB: db}
}

// GetAll fetches all categories in database.
func (r *PostgresCategoryRepository) GetAll() ([]*models.Category, error) {
	var categories []*models.Category
	stmt := `SELECT id, name, created_at FROM categories`
	query, err := r.DB.Prepare(stmt)
	if err != nil {
		return nil, err
	}
	defer handleError(query)

	rows, err := query.Query()
	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	for rows.Next() {
		var category models.Category
		err = rows.Scan(&category.ID, &category.Name, &category.CreatedAt)

		if err != nil {
			return nil, err
		}

		categories = append(categories, &category)
	}

	return categories, nil
}

// GetByID fetches a category with specified ID.
func (r *PostgresCategoryRepository) GetByID(categoryID string) (*models.Category, error) {
	category := new(models.Category)
	stmt := `SELECT id, name, created_at
FROM categories WHERE id = $1`

	query, err := r.DB.Prepare(stmt)
	if err != nil {
		return nil, err
	}
	defer handleError(query)

	err = query.QueryRow(categoryID).Scan(&category.ID, &category.Name, &category.CreatedAt)
	if err != nil {
		return nil, err
	}
	return category, nil
}

// GetByName finds category by category name.
func (r *PostgresCategoryRepository) GetByName(name string) (*models.Category, error) {
	category := new(models.Category)

	stmt := `SELECT id, name, created_at FROM categories WHERE name = $1`
	query, err := r.DB.Prepare(stmt)
	if err != nil {
		return nil, err
	}
	defer handleError(query)

	err = query.QueryRow(name).Scan(&category.ID, &category.Name, &category.CreatedAt)
	if err != nil {
		return nil, err
	}
	return category, nil
}

// Create creates new category
func (r *PostgresCategoryRepository) Create(category *models.Category) error {
	stmt := `INSERT INTO categories (name, created_at) VALUES($1,$2)`

	query, err := r.DB.Prepare(stmt)
	if err != nil {
		return err
	}
	defer handleError(query)

	_, err = query.Exec(category.Name, category.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

// Update updates a category by ID.
func (r *PostgresCategoryRepository) Update(categoryID string, category *models.Category) error {
	stmt := `UPDATE categories SET name = $2 WHERE id = $1`

	query, err := r.DB.Prepare(stmt)
	if err != nil {
		return err
	}
	defer handleError(query)

	_, err = query.Exec(categoryID, category.Name)
	if err != nil {
		return err
	}

	return nil
}

// Delete deletes a category by ID.
func (r *PostgresCategoryRepository) Delete(categoryID string) error {
	stmt := `DELETE FROM categories WHERE id = $1`
	query, err := r.DB.Prepare(stmt)
	if err != nil {
		return err
	}
	defer handleError(query)

	_, err = query.Exec(categoryID)
	if err != nil {
		return err
	}
	return nil
}
