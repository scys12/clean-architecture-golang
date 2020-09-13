package module

import (
	rRole "github.com/scys12/clean-architecture-golang/repository/role"
	rUser "github.com/scys12/clean-architecture-golang/repository/user"
	uUser "github.com/scys12/clean-architecture-golang/usecase/user"
)

type usecase struct {
	userRepo rUser.Repository
	roleRepo rRole.Repository
}

func New(uRepo rUser.Repository, rRepo rRole.Repository) uUser.Usecase {
	return &usecase{
		userRepo: uRepo,
		roleRepo: rRepo,
	}
}
