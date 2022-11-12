package service

import (
	"errors"
	"github.com/kirill0909/neurohacking-api/models"
	"github.com/kirill0909/neurohacking-api/pkg/repository"
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
	return w.repo.Create(word, userId, categoryId)
}

func (w *WordService) CheckCategoryOwner(userId, categoryId int) bool {
	return w.repo.CheckCategoryOwner(userId, categoryId)
}
