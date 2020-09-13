package database

import (
	"context"
	"fmt"

	"github.com/scys12/clean-architecture-golang/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	Client   *mongo.Client
	Context  context.Context
	Database *mongo.Database
}

func NewMongoDB(config *config.Config) (*MongoClient, error) {
	uri := fmt.Sprintf("%v://%v:%v/%v", config.DBDriver, config.DBHost, config.DBPort, config.DBName)
	clientOpt := options.Client().ApplyURI(uri)
	client, err := mongo.NewClient(clientOpt)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	mongoClient := &MongoClient{
		Client:   client,
		Context:  ctx,
		Database: client.Database(config.DBName),
	}
	return mongoClient, nil
}
