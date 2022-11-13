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

func (w *WordPostgres) GetAll(userId, categoryId int) ([]models.Word, error) {
	var words []models.Word

	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id=$1 AND category_id=$2", wordsTable)
	err := w.db.Select(&words, query, userId, categoryId)

	return words, err
}

func (w *WordPostgres) GetById(userId, categoryId, wordId int) (models.Word, error) {
	var word models.Word

	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id=$1 AND category_id=$2 AND id=$3", wordsTable)
	err := w.db.Get(&word, query, userId, categoryId, wordId)

	return word, err
}

func (w *WordPostgres) Update(input models.WordUpdateInput, userId, categoryId, wordId int) (models.Word, error) {
	var updatedWord models.Word

	query := fmt.Sprintf(`UPDATE %s SET name=$1, last_update=now() WHERE user_id=$2 AND category_id=$3 AND id=$4
	RETURNING id, user_id, category_id, name, date_creation, last_update`, wordsTable)

	row := w.db.QueryRow(query, input.Name, userId, categoryId, wordId)
	err := row.Scan(&updatedWord.Id, &updatedWord.UID, &updatedWord.CategoryId, &updatedWord.Name,
		&updatedWord.DateCreation, &updatedWord.LastUpdate)

	return updatedWord, err
}
