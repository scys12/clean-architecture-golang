package database

import (
	"context"
	"fmt"
	"time"

	"github.com/scys12/clean-architecture-golang/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	Client  *mongo.Client
	Context context.Context
}

func NewMongoDB(config *config.Config) (*MongoClient, error) {
	uri := fmt.Sprintf("%v://%v:%v", config.DBDriver, config.DBHost, config.DBPort)
	clientOpt := options.Client().ApplyURI(uri)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOpt)
	if err != nil {
		return nil, err
	}
	mongoClient := &MongoClient{
		Client:  client,
		Context: ctx,
	}
	return mongoClient, nil
}
