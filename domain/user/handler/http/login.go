package http

import (
	user_payload "gateway-grpc/domain/user/payload"
	user_response "gateway-grpc/domain/user/response"
	"gateway-grpc/lib/constant"
	"gateway-grpc/lib/helper"
	"gateway-grpc/lib/pkg/utils"

	"github.com/labstack/echo/v4"
)

// Login godoc
// @Summary login user
// @Description login user
// @Tags User
// @Accept  json
// @Produce  json
// @Param request body user_payload.LoginInput true "Request Body"
// @Success 200 {object} user_response.LoginSwagger
// @Failure 400 {object} user_response.BaseSwagger
// @Failure 409 {object} user_response.BaseSwagger
// @Failure 500 {object} user_response.BaseSwagger
// @Router /user/login [post]
func (h *UserHandler) Login(c echo.Context) error {
	var req user_payload.LoginInput
	c.Bind(&req)

	ctx, requestID := helper.GenerateRequestID(c)
	c.Response().Header().Set("X-Request-ID", requestID)

	login, err := h.user.Login(ctx, req.ToPB())
	if err != nil {
		return utils.ErrorResponse(c, err, map[string]interface{}{})
	}

	data := user_response.LoginProtoToJsonResponse(login)

	return utils.SuccessResponse(c, constant.SuccessGetData, data)
}
