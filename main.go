package main

import (
	"gateway-grpc/config"
	"log"
	"os"

	user_handler "gateway-grpc/domain/user/handler/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/grpc"
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
	serviceName := "USER_SERVICE"

	grpcAddress := os.Getenv(serviceName)
	grpcConn, err := grpc.Dial(grpcAddress)

	if err != nil {
		log.Fatal("did not connect: %s", err)
	} else {
		log.Println("connected to : " + grpcAddress)
	}

	user_handler.NewUserHandler(e, grpcConn)
	// e.GET("/swagger/*", echoSwagger.WrapHandler)

	// start serve
	port := "PORT"
	portNumber := os.Getenv(port)
	e.Logger.Fatal(e.Start(":" + portNumber))
}
