package helper

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/metadata"
)

func GenerateRequestID(c echo.Context) (ctx context.Context, requestID string) {

	requestID = uuid.New().String()

	md := metadata.Pairs("request_id", requestID)

	cekHeader := c.Request().Header

	tokenString := strings.Replace(cekHeader.Get("X-Authorization"), "JWT ", "", -1)

	userId, _ := GetDataFromToken(tokenString)

	if userId != "" {
		md.Append("user_id", userId)
	}

	return metadata.NewOutgoingContext(context.Background(), md), requestID
}
