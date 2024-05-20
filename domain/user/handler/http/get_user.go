package http

import (
	"gateway-grpc/lib/constant"
	"gateway-grpc/lib/helper"
	"gateway-grpc/lib/pkg/utils"

	user_response "gateway-grpc/domain/user/response"

	"github.com/labstack/echo/v4"
)

// Get User Profile godoc
// @Summary get user profile
// @Description get user profile
// @Tags User
// @Accept  json
// @Produce  json
// @Success 200 {object} user_response.UserSwagger
// @Failure 400 {object} user_response.BaseSwagger
// @Failure 403 {object} user_response.BaseSwagger
// @Failure 500 {object} user_response.BaseSwagger
// @Router /user/profile [get]
// @Param  X-Authorization header string true "X-Authorization" default(JWT <token-from-login>) example(JWT eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c)
func (h *UserHandler) GetUser(c echo.Context) error {

	ctx, requestID := helper.GenerateRequestID(c)
	c.Response().Header().Set("X-Request-ID", requestID)

	user, err := h.user.GetUser(ctx, nil)
	if err != nil {
		return utils.ErrorResponse(c, err, map[string]interface{}{})
	}

	data := user_response.UserProtoToJsonResponse(user)

	return utils.SuccessResponse(c, constant.SuccessGetData, data)
}
