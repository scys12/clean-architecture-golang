package module

import (
	rCategory "github.com/scys12/clean-architecture-golang/repository/category"
	uCategory "github.com/scys12/clean-architecture-golang/usecase/category"
)

type categoryUsecase struct {
	repo rCategory.Repository
}

func New(catRepo rCategory.Repository) uCategory.Usecase {
	return &categoryUsecase{
		repo: catRepo,
	}
}
