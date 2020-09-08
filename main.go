package main

import (
	"log"

	dUserHttp "github.com/scys12/clean-architecture-golang/delivery/user/http"

	uUserModule "github.com/scys12/clean-architecture-golang/usecase/user/module"

	rUser "github.com/scys12/clean-architecture-golang/repository/user/module"

	"github.com/scys12/clean-architecture-golang/delivery/middleware"
	"github.com/scys12/clean-architecture-golang/session"

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

	rd := session.NewRedisPool(config)
	rdConn := rd.Pool.Get()

	defer rdConn.Close()
	if err := session.Ping(rdConn); err != nil {
		panic(err)
	}
	rd.Conn = rdConn

	defer func() {
		if err = client.Client.Disconnect(client.Context); err != nil {
			panic(err)
		}
	}()
	if err := client.Client.Ping(client.Context, readpref.Primary()); err != nil {
		panic(err)
	}

	e := echo.New()
	e.Use(middleware.CORS())

	categoryRepo := rCategory.New(client.Database)
	categoryUC := uCategoryModule.New(categoryRepo)
	dCategoryHandler := dCategoryHttp.New(categoryUC)
	dCategoryHttp.SetRoute(e, dCategoryHandler)

	userRepo := rUser.New(client.Database)
	userUC := uUserModule.New(userRepo)
	dUserHandler := dUserHttp.New(userUC, rd)
	dUserHttp.SetRoute(e, dUserHandler, rd)

	log.Fatal(e.Start(":8080"))
}
