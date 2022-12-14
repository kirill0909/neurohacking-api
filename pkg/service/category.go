package service

import (
	"errors"
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

func (c *CategoryService) GetById(userId, categoryId int) (models.Category, error) {
	return c.repo.GetById(userId, categoryId)
}

func (c *CategoryService) Update(input models.CategoryUpdateInput, userId, categoryId int) (models.Category, error) {
	if ok := c.repo.CheckCategoryIdExists(userId, categoryId); !ok {
		return models.Category{}, errors.New("user does not have this category")
	}
	return c.repo.Update(input, userId, categoryId)
}

func (c *CategoryService) Delete(userId, categoryId int) (models.Category, error) {
	if ok := c.repo.CheckCategoryIdExists(userId, categoryId); !ok {
		return models.Category{}, errors.New("user does not have this category")
	}
	return c.repo.Delete(userId, categoryId)
}

func (c *CategoryService) CheckCategoryIdExists(userId, categoryId int) bool {
	return c.repo.CheckCategoryIdExists(userId, categoryId)
}
