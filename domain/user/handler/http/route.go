package http

import (
	userpb "gateway-grpc/pb/user"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	user userpb.UserServiceClient
}

func NewUserHandler(e *echo.Echo, user userpb.UserServiceClient) {
	handler := &UserHandler{
		user: user,
	}

	e.POST("/user/register", handler.RegisterUser)
	e.POST("/user/login", handler.Login)
	e.POST("/user/token", handler.GetToken)
}
