package module

import (
	"context"

	"github.com/scys12/clean-architecture-golang/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *repository) GetAllCategories(ctx context.Context) (cats []*model.Category, err error) {
	cur, err := r.db.Collection(r.collection).Find(ctx, bson.M{})
	if err != nil {
		return
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var cat *model.Category
		err := cur.Decode(&cat)
		if err != nil {
			return nil, err
		}
		cats = append(cats, cat)
	}
	return
}
