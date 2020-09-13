package module

import (
	"context"

	"github.com/scys12/clean-architecture-golang/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *repository) GetRoleByName(ctx context.Context, roleName string) (*model.Role, error) {
	role := new(model.Role)
	err := r.db.Collection(r.collection).FindOne(ctx, bson.M{"name": roleName}).Decode(&role)
	if err != nil {
		return nil, err
	}
	return role, nil
}
