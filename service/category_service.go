package service

import (
	"errors"
	"golangunittest/entity"
	"golangunittest/repository"
)

type CategoryService struct {
	CategoryRepository repository.CategoryRepository
}

func (categoryService *CategoryService) FindByCategoryID(categoryID int) (*entity.Category, error) {

	category, _ := categoryService.CategoryRepository.FindByID(categoryID)

	if category == nil {
		return category, errors.New("category is'nt found")
	}

	return category, nil

}
