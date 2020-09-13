package category

import (
	"context"

	"github.com/scys12/clean-architecture-golang/model"

	"github.com/scys12/clean-architecture-golang/pkg/payload/request"
)

type Usecase interface {
	AuthenticateUser(context.Context, *request.LoginRequest) (*model.User, error)
	RegisterUser(context.Context, *request.RegisterRequest) error
	EditUserProfile(context.Context) error
}
