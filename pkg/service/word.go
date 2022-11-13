package service

import (
	"errors"
	"github.com/kirill0909/neurohacking-api/models"
	"github.com/kirill0909/neurohacking-api/pkg/repository"
	"strings"
)

type WordService struct {
	repo repository.Word
}

func NewWordService(repo repository.Word) *WordService {
	return &WordService{repo: repo}
}

func (w *WordService) Create(word models.Word, userId, categoryId int) (models.Word, error) {
	if ok := w.repo.CheckCategoryOwner(userId, categoryId); !ok {
		return models.Word{}, errors.New("This user is not the owner of this category")
	}
	word.Name = strings.Title(strings.ToLower(word.Name))
	return w.repo.Create(word, userId, categoryId)
}

func (w *WordService) CheckCategoryOwner(userId, categoryId int) bool {
	return w.repo.CheckCategoryOwner(userId, categoryId)
}

func (w *WordService) GetAll(userId, categoryId int) ([]models.Word, error) {
	if ok := w.repo.CheckCategoryOwner(userId, categoryId); !ok {
		return []models.Word{}, errors.New("This user is not the owner of this category")
	}
	return w.repo.GetAll(userId, categoryId)
}

func (w *WordService) GetById(userId, categoryId, wordId int) (models.Word, error) {
	if ok := w.repo.CheckCategoryOwner(userId, categoryId); !ok {
		return models.Word{}, errors.New("This user is not the owner of this category")
	}
	return w.repo.GetById(userId, categoryId, wordId)
}

func (w *WordService) Update(input models.WordUpdateInput, userId, categoryId, wordId int) (models.Word, error) {
	if ok := w.repo.CheckCategoryOwner(userId, categoryId); !ok {
		return models.Word{}, errors.New("This user is not the owner of this category")
	}

	return w.repo.Update(input, userId, categoryId, wordId)
}
