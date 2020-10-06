package main

import (
	"log"

	"github.com/scys12/clean-architecture-golang/pkg/validator"

	"github.com/scys12/clean-architecture-golang/pkg/aws"

	dCategoryHttp "github.com/scys12/clean-architecture-golang/delivery/category/http"
	dItemHttp "github.com/scys12/clean-architecture-golang/delivery/item/http"
	dUserHttp "github.com/scys12/clean-architecture-golang/delivery/user/http"

	uCategoryModule "github.com/scys12/clean-architecture-golang/usecase/category/module"
	uItemModule "github.com/scys12/clean-architecture-golang/usecase/item/module"
	uUserModule "github.com/scys12/clean-architecture-golang/usecase/user/module"

	rCategory "github.com/scys12/clean-architecture-golang/repository/category/module"
	rItem "github.com/scys12/clean-architecture-golang/repository/item/module"
	rRole "github.com/scys12/clean-architecture-golang/repository/role/module"
	rUser "github.com/scys12/clean-architecture-golang/repository/user/module"

	"github.com/scys12/clean-architecture-golang/delivery/middleware"
	"github.com/scys12/clean-architecture-golang/pkg/session"

	"github.com/labstack/echo/v4"

	"github.com/scys12/clean-architecture-golang/pkg/database"

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
	roleRepo := rRole.New(mongo.Database)
	userUC := uUserModule.New(userRepo, roleRepo, aws_s3, rd)
	dUserHandler := dUserHttp.New(userUC)
	dUserHttp.SetRoute(e, dUserHandler, rd)

	itemRepo := rItem.New(mongo.Database)
	itemUC := uItemModule.New(itemRepo, aws_s3, rd)
	dItemHandler := dItemHttp.New(itemUC)
	dItemHttp.SetRoute(e, dItemHandler, rd)

	log.Fatal(e.Start(":8080"))
}
