package service

import (
	"github.com/kirill0909/neurohacking-api/pkg/repository"
)

type User interface{}

type Category interface{}

type Word interface{}

type Service struct {
	User
	Category
	Word
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
