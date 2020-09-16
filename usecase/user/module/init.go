package module

import (
	awsS3 "github.com/scys12/clean-architecture-golang/pkg/aws"
	rRole "github.com/scys12/clean-architecture-golang/repository/role"
	rUser "github.com/scys12/clean-architecture-golang/repository/user"
	uUser "github.com/scys12/clean-architecture-golang/usecase/user"
)

type usecase struct {
	userRepo rUser.Repository
	roleRepo rRole.Repository
	awsStore awsS3.S3Store
}

func New(uRepo rUser.Repository, rRepo rRole.Repository, awsS3Store awsS3.S3Store) uUser.Usecase {
	return &usecase{
		userRepo: uRepo,
		roleRepo: rRepo,
		awsStore: awsS3Store,
	}
}
