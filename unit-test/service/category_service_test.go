package service

import (
	"testing"
	"unit-test/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = repository.CategoryRepositoryMock{Mock: mock.Mock{}}

func TestCategoryService_Get(t *testing.T) {
	// program mock
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	categoryService := CategoryService{
		Repository: &categoryRepository,
	}

	category, err := categoryService.Get("1")

	assert.Nil(t, category)
	assert.NotNil(t, err)
}
