package module

import (
	awsS3 "github.com/scys12/clean-architecture-golang/pkg/aws"
	rItem "github.com/scys12/clean-architecture-golang/repository/item"
	uItem "github.com/scys12/clean-architecture-golang/usecase/item"
)

type usecase struct {
	itemRepo rItem.Repository
	awsStore awsS3.S3Store
}

func New(itemRepo rItem.Repository, awsS3Store awsS3.S3Store) uItem.Usecase {
	return &usecase{
		itemRepo: itemRepo,
		awsStore: awsS3Store,
	}
}
