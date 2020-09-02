package category

import (
	"github.com/scys12/clean-architecture-golang/models"
)

type Repository interface {
	GetAllCategories() []models.Category
}
