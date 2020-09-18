package module

import (
	"github.com/scys12/clean-architecture-golang/repository/item"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	db         *mongo.Database
	collection string
}

func New(db *mongo.Database) item.Repository {
	return &repository{
		db:         db,
		collection: "item",
	}
}
