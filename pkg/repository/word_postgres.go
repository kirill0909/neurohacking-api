package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/kirill0909/neurohacking-api/models"
)

type WordPostgres struct {
	db *sqlx.DB
}

func NewWordPostgres(db *sqlx.DB) *WordPostgres {
	return &WordPostgres{db: db}
}

func (w *WordPostgres) Create(word models.Word, userId, categoryId int) (models.Word, error) {
	var insertedWord models.Word

	query := fmt.Sprintf(`INSERT INTO %s (user_id, category_id, name, date_creation, last_update)
	VALUES ($1, $2, $3, now(), now()) RETURNING id, user_id, category_id, name, date_creation, last_update`, wordsTable)

	row := w.db.QueryRow(query, userId, categoryId, word.Name)
	err := row.Scan(&insertedWord.Id, &insertedWord.UID, &insertedWord.CategoryId, &insertedWord.Name, &insertedWord.DateCreation,
		&insertedWord.LastUpdate)

	return insertedWord, err
}

func (w *WordPostgres) CheckCategoryOwner(userId, categoryId int) bool {
	var result bool

	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE user_id=$1 AND id=$2)", categoriesTable)
	row := w.db.QueryRow(query, userId, categoryId)
	row.Scan(&result)

	return result
}
