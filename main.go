package main

import (
	"gateway-grpc/config"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	_, err := config.Setup()
	if err != nil {
		log.Fatal("Cannot load config ", err.Error())
	}

	e := echo.New()

	//middleware
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())

	//depedency
	// authRepo := repo.NewAuthRepo(db)

	// authUc := usecase.NewAuthUsecase(authRepo)

	// handler.NewAuthHander(e, authUc)
	// e.GET("/swagger/*", echoSwagger.WrapHandler)

	// start serve
	port := "PORT"
	portNumber := os.Getenv(port)
	e.Logger.Fatal(e.Start(portNumber))
}
