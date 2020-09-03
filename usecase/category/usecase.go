package category

import (
	"context"

	models "github.com/scys12/clean-architecture-golang/model"
)

type Usecase interface {
	GetAllCategories(context.Context) ([]*models.Category, error)
}
