package module

import (
	"context"
	"time"

	"github.com/scys12/clean-architecture-golang/repository/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type repository struct {
	db         *mongo.Database
	collection string
}

const coll = "user"

func New(db *mongo.Database) user.Repository {
	createIndexes(db)
	return &repository{
		db:         db,
		collection: coll,
	}
}

func createIndexes(db *mongo.Database) {
	models := []mongo.IndexModel{
		{
			Keys:    bson.D{{"username", 1}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys:    bson.D{{"email", 1}},
			Options: options.Index().SetUnique(true),
		},
	}
	opts := options.CreateIndexes().SetMaxTime(2 * time.Second)
	_, err := db.Collection(coll).Indexes().CreateMany(context.TODO(), models, opts)
	if err != nil {
		panic(err)
	}
}
