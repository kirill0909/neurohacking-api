package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/kirill0909/neurohacking-api/models"
)

type CategoryPostgres struct {
	db *sqlx.DB
}

func NewCategoryPostgres(db *sqlx.DB) *CategoryPostgres {
	return &CategoryPostgres{db: db}
}

func (c *CategoryPostgres) Create(category models.Category, userId int) (models.Category, error) {
	var insertedCategory models.Category

	query := fmt.Sprintf(`INSERT INTO %s (user_id, name, date_creation, last_update) 
	VALUES($1, $2, now(), now()) RETURNING id, user_id, name, date_creation, last_update`, categoriesTable)

	row := c.db.QueryRow(query, userId, category.Name)
	err := row.Scan(&insertedCategory.Id, &insertedCategory.UID, &insertedCategory.Name,
		&insertedCategory.DateCreation, &insertedCategory.LastUpdate)

	return insertedCategory, err
}

func (c *CategoryPostgres) GetAll(userId int) ([]models.Category, error) {
	var categories []models.Category

	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1", categoriesTable)
	err := c.db.Select(&categories, query, userId)

	return categories, err
}

func (c *CategoryPostgres) GetById(userId, categoryId int) (models.Category, error) {
	var category models.Category

	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id=$1 AND id=$2", categoriesTable)
	err := c.db.Get(&category, query, userId, categoryId)

	return category, err
}
