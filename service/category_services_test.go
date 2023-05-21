package service

import (
	"testing"

	"github.com/nisfu-saaban/mock/entity"
	"github.com/nisfu-saaban/mock/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_get(t *testing.T) {
	categoryRepository.Mock.On("FindByID", "1").Return(nil)
	catefory, err := categoryService.Get("1")
	assert.Nil(t, catefory)
	assert.NotNil(t, err)
}

func TestCategoryService_getFound(t *testing.T) {
	category := entity.Category{
		Id:   "2",
		Name: "Phone",
	}

	categoryRepository.Mock.On("FindByID", "2").Return(category)
	result, err := categoryService.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.Id, result.Id)
	assert.Equal(t, category.Name, result.Name)

}
