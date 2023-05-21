package service

import (
	"errors"

	"github.com/nisfu-saaban/mock/entity"
	"github.com/nisfu-saaban/mock/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindByID(id)
	if category == nil {
		return category, errors.New("Category not found")
	} else {
		return category, nil
	}
}
