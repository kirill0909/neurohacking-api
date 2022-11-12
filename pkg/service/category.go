package service

import (
	"github.com/kirill0909/neurohacking-api/models"
	"github.com/kirill0909/neurohacking-api/pkg/repository"
	"strings"
)

type CategoryService struct {
	repo repository.Category
}

func NewCategoryService(repo repository.Category) *CategoryService {
	return &CategoryService{repo: repo}
}

func (c *CategoryService) Create(category models.Category, userId int) (models.Category, error) {
	category.Name = strings.Title(strings.ToLower(category.Name))
	return c.repo.Create(category, userId)
}

func (c *CategoryService) GetAll(userId int) ([]models.Category, error) {
	return c.repo.GetAll(userId)
}
