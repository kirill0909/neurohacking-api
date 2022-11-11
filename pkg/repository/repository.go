package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/kirill0909/neurohacking-api/models"
)

type User interface {
	Create(user models.User) (int, error)
	GetUser(email, password string) (models.User, error)
	CheckUserIdExists(id int) (bool, error)
	Update(input models.UserUpdateInput, userId int) error
	Delete(userId int) error
}

type Category interface {
	Create(category models.Category, userId int) (models.Category, error)
}

type Word interface{}

type Repository struct {
	User
	Category
	Word
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User:     NewUserPostgres(db),
		Category: NewCategoryPostgres(db),
	}
}
