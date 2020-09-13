package module

import (
	"context"

	"github.com/scys12/clean-architecture-golang/pkg/payload/request"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/scys12/clean-architecture-golang/model"
)

func (r *repository) GetUserAuthenticateData(ctx context.Context, uName string) (user *model.User, err error) {
	err = r.db.Collection(r.collection).FindOne(ctx, bson.M{"username": uName}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *repository) RegisterUser(ctx context.Context, req *request.RegisterRequest) (err error) {
	data := bson.M{
		"username": req.Username,
		"password": req.Password,
		"email":    req.Email,
		"name":     req.Name,
		"location": req.Location,
		"phone":    req.Phone,
		"role":     req.Role,
	}
	_, err = r.db.Collection(r.collection).InsertOne(ctx, data)
	return
}
