package main

import (
	"log"

	"github.com/labstack/echo/v4"

	"github.com/scys12/clean-architecture-golang/database"
	dCategoryHttp "github.com/scys12/clean-architecture-golang/delivery/category/http"
	rCategory "github.com/scys12/clean-architecture-golang/repository/category/module"
	uCategoryModule "github.com/scys12/clean-architecture-golang/usecase/category/module"

	"github.com/scys12/clean-architecture-golang/config"

	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	config := config.NewConfig()

	client, err := database.NewMongoDB(config)

	defer func() {
		if err = client.Client.Disconnect(client.Context); err != nil {
			panic(err)
		}
	}()
	if err := client.Client.Ping(client.Context, readpref.Primary()); err != nil {
		panic(err)
	}

	e := echo.New()
	categoryRepo := rCategory.New(client.Database)
	categoryUC := uCategoryModule.New(categoryRepo)
	dCategoryHandler := dCategoryHttp.New(categoryUC)
	dCategoryHttp.SetRoute(e, dCategoryHandler)
	log.Fatal(e.Start(":8080"))
}
