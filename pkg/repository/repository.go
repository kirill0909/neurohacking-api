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
	GetAll(userId int) ([]models.Category, error)
	GetById(userId, categoryId int) (models.Category, error)
	CheckCategoryIdExists(userId, categoryId int) bool
	Update(input models.CategoryUpdateInput, userId, categoryId int) (models.Category, error)
	Delete(userId, categoryId int) (models.Category, error)
}

type Word interface {
	Create(word models.Word, userId, categoryId int) (models.Word, error)
	CheckCategoryOwner(userId, categoryId int) bool
}

type Repository struct {
	User
	Category
	Word
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User:     NewUserPostgres(db),
		Category: NewCategoryPostgres(db),
		Word:     NewWordPostgres(db),
	}
}
