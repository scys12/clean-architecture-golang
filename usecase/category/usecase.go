package category

import (
	"github.com/scys12/clean-architecture-golang/models"
)

type Usecase interface {
	GetAllCategories() []models.Category
}
