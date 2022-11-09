package repository

import (
	"github.com/jmoiron/sqlx"
)

type User interface{}

type Category interface{}

type Word interface{}

type Repository struct {
	User
	Category
	Word
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
