package main

import (
	"log"

	"github.com/scys12/clean-architecture-golang/pkg/validator"

	"github.com/scys12/clean-architecture-golang/pkg/aws"

	dUserHttp "github.com/scys12/clean-architecture-golang/delivery/user/http"

	uUserModule "github.com/scys12/clean-architecture-golang/usecase/user/module"

	rRole "github.com/scys12/clean-architecture-golang/repository/role/module"
	rUser "github.com/scys12/clean-architecture-golang/repository/user/module"

	"github.com/scys12/clean-architecture-golang/delivery/middleware"
	"github.com/scys12/clean-architecture-golang/pkg/session"

	"github.com/labstack/echo/v4"

	dCategoryHttp "github.com/scys12/clean-architecture-golang/delivery/category/http"
	"github.com/scys12/clean-architecture-golang/pkg/database"
	rCategory "github.com/scys12/clean-architecture-golang/repository/category/module"
	uCategoryModule "github.com/scys12/clean-architecture-golang/usecase/category/module"

	"github.com/scys12/clean-architecture-golang/pkg/config"

	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	config := config.NewConfig()
	aws_s3, err := aws.InitializeSessionAWS()
	if err != nil {
		panic(err)
	}
	mongo, err := database.NewMongoDB(config)

	rd := session.NewRedisPool(config)
	defer rd.Connect().Close()
	if err := session.Ping(rd); err != nil {
		panic(err)
	}
	defer func() {
		if err = mongo.Client.Disconnect(mongo.Context); err != nil {
			panic(err)
		}
	}()
	if err := mongo.Client.Ping(mongo.Context, readpref.Primary()); err != nil {
		panic(err)
	}

	e := echo.New()
	e.Use(middleware.CORS())
	e.Validator = validator.New()
	categoryRepo := rCategory.New(mongo.Database)
	categoryUC := uCategoryModule.New(categoryRepo)
	dCategoryHandler := dCategoryHttp.New(categoryUC)
	dCategoryHttp.SetRoute(e, dCategoryHandler)

	userRepo := rUser.New(mongo.Database)
	mockRepo := rRole.New(mongo.Database)
	userUC := uUserModule.New(userRepo, mockRepo, aws_s3)
	dUserHandler := dUserHttp.New(userUC, rd)
	dUserHttp.SetRoute(e, dUserHandler, rd)

	log.Fatal(e.Start(":8080"))
}
