package service

import (
	"github.com/kirill0909/neurohacking-api/models"
	"github.com/kirill0909/neurohacking-api/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type User interface {
	Create(user models.User) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, error)
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

type Word interface{}

type Service struct {
	User
	Category
	Word
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User:     NewUserService(repos.User),
		Category: NewCategoryService(repos.Category),
	}
}
