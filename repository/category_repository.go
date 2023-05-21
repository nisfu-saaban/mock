package repository

import "github.com/nisfu-saaban/mock/entity"

type CategoryRepository interface {
	FindByID(id string) *entity.Category
}
