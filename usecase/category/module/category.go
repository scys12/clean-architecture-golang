package module

import (
	"github.com/scys12/clean-architecture-golang/models"
)

func (u *categoryUsecase) GetAllCategories() []models.Category {
	categories := u.repo.GetAllCategories()
}
