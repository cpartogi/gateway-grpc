package main

import (
	"gateway-grpc/config"
	"log"
	"os"

	user_handler "gateway-grpc/domain/user/handler/http"

	userpb "gateway-grpc/pb/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	userServiceName := "USER_SERVICE"

	userGrpcAddress := os.Getenv(userServiceName)
	userGrpcConn, err := grpc.Dial(userGrpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Println("failed to connect user service " + err.Error())
	} else {
		log.Println("connected to user service at : " + userGrpcAddress)
	}
	userGrpcService := userpb.NewUserServiceClient(userGrpcConn)

	user_handler.NewUserHandler(e, userGrpcService)
	// e.GET("/swagger/*", echoSwagger.WrapHandler)

	// start serve
	port := "PORT"
	portNumber := os.Getenv(port)
	e.Logger.Fatal(e.Start(":" + portNumber))
}
