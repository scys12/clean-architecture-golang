package module

import (
	"context"

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

func (r *repository) RegisterUser(ctx context.Context, user model.User) (err error) {
	data := bson.M{
		"username": user.Username,
		"password": user.Password,
		"email":    user.Email,
		"name":     user.Name,
		"location": user.Location,
		"phone":    user.Phone,
		"role":     user.Role,
	}
	_, err = r.db.Collection(r.collection).InsertOne(ctx, data)
	return
}
