package route

import (
	"context"
	"log"
	"os"

	user_handler "gateway-grpc/domain/user/handler/http"
	userpb "gateway-grpc/pb/user"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func SetupRouterUser(ctx context.Context, e *echo.Echo) {
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
}
