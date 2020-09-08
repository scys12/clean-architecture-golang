package module

import (
	"github.com/scys12/clean-architecture-golang/repository/user"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	db         *mongo.Database
	collection string
}

func New(db *mongo.Database) user.Repository {
	return &repository{
		db:         db,
		collection: "user",
	}
}
