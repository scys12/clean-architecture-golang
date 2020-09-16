package user

import (
	"context"

	"github.com/scys12/clean-architecture-golang/model"
)

type Repository interface {
	GetUserAuthenticateData(context.Context, map[string]interface{}) (*model.UserAuth, *model.UserProfile, error)
	RegisterUser(context.Context, model.UserAuth) error
	EditUserProfile(context.Context, model.UserProfile) error
}
