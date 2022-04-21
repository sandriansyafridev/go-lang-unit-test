package service

import (
	"errors"
	"golangunittest/entity"
	"golangunittest/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepositoryMock = repository.CategoryRepositoryMock{
	Mock: mock.Mock{},
}

var categoryService = CategoryService{
	CategoryRepository: &categoryRepositoryMock,
}

func TestNotFound_FindByCategoryID(t *testing.T) {
	category := entity.Category{}
	err := errors.New("not found category")
	categoryRepositoryMock.Mock.On("FindByID", 1).Return(category, err)
	result, err := categoryService.FindByCategoryID(1)

	assert.Nil(t, err)
	assert.Equal(t, category.ID, result.ID)
	assert.Equal(t, category.Name, result.Name)

}

func TestFound_FindByCategoryID(t *testing.T) {

	category := entity.Category{
		ID:   2,
		Name: "Sport",
	}
	categoryRepositoryMock.Mock.On("FindByID", 2).Return(category, nil)

	result, err := categoryService.FindByCategoryID(2)

	assert.Nil(t, err)
	assert.NotNil(t, result)

	assert.Equal(t, category.ID, result.ID)
	assert.Equal(t, category.Name, result.Name)

}
