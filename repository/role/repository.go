package role

import (
	"context"

	"github.com/scys12/clean-architecture-golang/model"
)

type Repository interface {
	GetRoleByName(context.Context, string) (*model.Role, error)
}
