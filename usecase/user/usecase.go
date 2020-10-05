package user

import (
	"context"

	"github.com/scys12/clean-architecture-golang/pkg/payload/request"
)

type Usecase interface {
	AuthenticateUser(context.Context, *request.LoginRequest) (*AuthenticateResponse, error)
	RegisterUser(context.Context, *request.RegisterRequest) error
	EditUserProfile(context.Context, *request.ProfileRequest) (*Response, error)
	GetUserProfile(context.Context, string) (*Response, error)
	Logout(context.Context, string, string) error
}
