package repository

import "golangunittest/entity"

type CategoryRepository interface {
	FindByID(categoryID int) (*entity.Category, error)
}
