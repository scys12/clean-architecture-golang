package main

import (
	"context"
	"fmt"
	"time"

	"github.com/scys12/clean-architecture-golang/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	config := config.NewConfig()

	uri := fmt.Sprintf("%v://%v:%v", config.DBDriver, config.DBHost, config.DBPort)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = mongoClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	if err := mongoClient.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

}
