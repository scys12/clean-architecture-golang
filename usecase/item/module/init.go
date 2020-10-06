package module

import (
	awsS3 "github.com/scys12/clean-architecture-golang/pkg/aws"
	"github.com/scys12/clean-architecture-golang/pkg/session"
	rItem "github.com/scys12/clean-architecture-golang/repository/item"
	uItem "github.com/scys12/clean-architecture-golang/usecase/item"
)

type usecase struct {
	itemRepo rItem.Repository
	awsStore awsS3.S3Store
	inMemory session.SessionStore
}

func New(itemRepo rItem.Repository, awsS3Store awsS3.S3Store, inMemory session.SessionStore) uItem.Usecase {
	return &usecase{
		itemRepo: itemRepo,
		awsStore: awsS3Store,
		inMemory: inMemory,
	}
}
