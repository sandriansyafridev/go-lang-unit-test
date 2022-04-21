package repository

import (
	"golangunittest/entity"

	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	mock.Mock
}

func (r *CategoryRepositoryMock) FindByID(categoryID int) (*entity.Category, error) {

	args := r.Mock.Called(categoryID)

	result := args.Get(0).(entity.Category)
	return &result, nil

}
