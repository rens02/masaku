package main

import (
	"fmt"
	"masaku/config"
	"masaku/controller"
	"masaku/database"
	"masaku/helpers"
	"masaku/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	var config = config.InitConfig()

	db := database.InitModel(*config)
	database.Migrate(db)

	jwtInterface := helpers.New(config.Secret)
	openAiInterface := helpers.NewOpenAi(config.OpenAiKey)
	// CloudinaryInterface := helpers.NewCloudninary(config.CloudinaryKey)

	userControll := controller.NewUsersControl(db, jwtInterface)
	ResepControll := controller.NewResepControl(db)
	KategoriControll := controller.NewKategoriControl(db)
	GenereteControll := controller.NewGenerateControl(db, jwtInterface, openAiInterface)


	e.Pre(middleware.RemoveTrailingSlash())

	// Middleware CORS
	e.Use(middleware.CORS())

	// Middleware Logger
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, time=${time_rfc3339}\n",
	}))

	// Rute untuk pengguna
	routes.RouteUser(e, userControll, ResepControll, KategoriControll , GenereteControll,*config)

	// Jalankan server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.ServerPort)))
}