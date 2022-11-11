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

type Category interface{}

type Word interface{}

type Service struct {
	User
	Category
	Word
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repos.User),
	}
}
