package module

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/scys12/clean-architecture-golang/model"
)

func (r *repository) CreateItem(ctx context.Context, item model.Item) (err error) {
	_, err = r.db.Collection(r.collection).InsertOne(ctx, item)
	return
}
func (r *repository) UpdateItem(ctx context.Context, item model.Item) (err error) {
	filter := bson.D{primitive.E{Key: "_id", Value: item.ID}}
	update := bson.D{primitive.E{Key: "$set", Value: item}}
	err = r.db.Collection(r.collection).FindOneAndUpdate(ctx, filter, update).Err()
	return
}
func (r *repository) RemoveItem(ctx context.Context, itemID primitive.ObjectID) (err error) {
	filter := bson.D{primitive.E{Key: "_id", Value: itemID}}
	_, err = r.db.Collection(r.collection).DeleteOne(ctx, filter)
	return
}

func (r *repository) GetItemData(ctx context.Context, filter map[string]interface{}) (item *model.Item, err error) {
	err = r.db.Collection(r.collection).FindOne(ctx, filter).Decode(&item)
	return
}
