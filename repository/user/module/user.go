package module

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/scys12/clean-architecture-golang/model"
)

func (r *repository) GetUserAuthenticateData(ctx context.Context, filter map[string]interface{}) (userAuth *model.UserAuth, userProfile *model.UserProfile, err error) {
	err = r.db.Collection(r.collection).FindOne(ctx, filter).Decode(&userAuth)
	if err != nil {
		return nil, nil, err
	}
	err = r.db.Collection(r.collection).FindOne(ctx, bson.M{"_id": userAuth.ID}).Decode(&userProfile)
	if err != nil {
		return nil, nil, err
	}
	return
}

func (r *repository) RegisterUser(ctx context.Context, user model.UserAuth) (err error) {
	_, err = r.db.Collection(r.collection).InsertOne(ctx, user)
	return
}

func (r *repository) EditUserProfile(ctx context.Context, user model.UserProfile) (err error) {
	filter := bson.D{primitive.E{Key: "_id", Value: user.ID}}
	update := bson.D{primitive.E{Key: "$set", Value: user}}
	err = r.db.Collection(r.collection).FindOneAndUpdate(ctx, filter, update).Err()
	return
}
