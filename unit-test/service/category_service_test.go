package service

import (
	"testing"
	"unit-test/entity"
	"unit-test/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = repository.CategoryRepositoryMock{Mock: mock.Mock{}}

func TestCategoryService_GetFound(t *testing.T) {
	// program mock
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	categoryService := CategoryService{
		Repository: &categoryRepository,
	}

	category, err := categoryService.Get("1")

	assert.Nil(t, category)
	assert.NotNil(t, err)
}

func TestCategoryService_GetSuccess(t *testing.T) {
	category := entity.Category{
		Id:   "1",
		Name: "Laptop",
	}

	// program mock
	categoryRepository.Mock.On("FindById", "2").Return(category)

	categoryService := CategoryService{
		Repository: &categoryRepository,
	}

	result, err := categoryService.Get("2")

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.Id, result.Id)
	assert.Equal(t, category.Name, result.Name)
}
