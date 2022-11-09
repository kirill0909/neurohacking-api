package service

import (
	"crypto/sha256"
	"fmt"
	"github.com/kirill0909/neurohacking-api/models"
	"github.com/kirill0909/neurohacking-api/pkg/repository"
	"os"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) CreateUser(user models.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return u.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("SOLT"))))
}
