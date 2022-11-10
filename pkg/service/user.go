package service

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/kirill0909/neurohacking-api/models"
	"github.com/kirill0909/neurohacking-api/pkg/repository"
	"os"
	"time"
)

const (
	tokenTTL = time.Hour * 12
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func (u *UserService) CreateUser(user models.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return u.repo.CreateUser(user)
}

func (u *UserService) GenerateToken(email, password string) (string, error) {
	user, err := u.repo.GetUser(email, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			// The token will depricate in twelve hours
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			// The token was created at ...
			IssuedAt: time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(os.Getenv("SIGNATURE_KEY")))
}

func (u *UserService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(os.Getenv("SIGNATURE_KEY")), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of the type *tokenClaims")
	}

	exists, err := u.CheckUserIdExists(claims.UserId)
	if err != nil {
		return 0, err
	}
	if !exists {
		return 0, errors.New("user does not exists")
	}

	return claims.UserId, nil

}

func (u *UserService) CheckUserIdExists(id int) (bool, error) {
	return u.repo.CheckUserIdExists(id)
}

func (u *UserService) Update(input models.UserUpdateInput, id int) error {
	if input.Password != nil {
		*input.Password = generatePasswordHash(*input.Password)
	}
	return u.repo.Update(input, id)
}

func generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("SOLT"))))
}
