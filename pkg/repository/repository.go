package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/kirill0909/neurohacking-api/models"
)

type User interface {
	CreateUser(user models.User) (int, error)
}

type Category interface{}

type Word interface{}

type Repository struct {
	User
	Category
	Word
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: NewUserPostgres(db),
	}
}
