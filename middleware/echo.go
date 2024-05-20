package middleware

import (
	"gateway-grpc/lib/constant"
	"gateway-grpc/lib/helper"
	"gateway-grpc/lib/pkg/utils"
	"strings"

	"github.com/labstack/echo/v4"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		cekHeader := c.Request().Header

		if cekHeader.Get("X-Authorization") == "" {
			return utils.ErrorBadRequest(c, constant.ErrBadRequest, map[string]interface{}{})
		}

		if !strings.Contains(cekHeader.Get("X-Authorization"), "JWT") {
			return utils.ErrorBadRequest(c, constant.ErrBadRequest, map[string]interface{}{})
		}

		tokenString := strings.Replace(cekHeader.Get("X-Authorization"), "JWT ", "", -1)

		_, err := helper.GetDataFromToken(tokenString)

		if err != nil {
			return utils.ErrorBadRequest(c, constant.ErrTokenExpired, map[string]interface{}{})
		}

		return next(c)
	}
}
