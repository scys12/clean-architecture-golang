package user

import (
	"context"

	"github.com/scys12/clean-architecture-golang/model"
)

type Repository interface {
	GetUserAuthenticateData(context.Context, string) (*model.User, error)
	RegisterUser(context.Context, model.User) error
}
