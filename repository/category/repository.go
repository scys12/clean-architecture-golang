package category

import (
	"context"

	"github.com/scys12/clean-architecture-golang/model"
)

type Repository interface {
	GetAllCategories(context.Context) ([]*model.Category, error)
}
