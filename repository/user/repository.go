package user

import (
	"context"

	"github.com/scys12/clean-architecture-golang/payload/request"

	"github.com/scys12/clean-architecture-golang/model"
)

type Repository interface {
	GetUserAuthenticateData(context.Context, string) (*model.User, error)
	RegisterUser(context.Context, *request.RegisterRequest) error
}
