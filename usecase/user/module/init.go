package module

import (
	rUser "github.com/scys12/clean-architecture-golang/repository/user"
	uUser "github.com/scys12/clean-architecture-golang/usecase/user"
)

type usecase struct {
	repo rUser.Repository
}

func New(catRepo rUser.Repository) uUser.Usecase {
	return &usecase{
		repo: catRepo,
	}
}
