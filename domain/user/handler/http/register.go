package http

import (
	user_payload "gateway-grpc/domain/user/payload"
	"gateway-grpc/lib/constant"
	"gateway-grpc/lib/helper"
	"gateway-grpc/lib/pkg/utils"

	user_response "gateway-grpc/domain/user/response"

	"github.com/labstack/echo/v4"
)

// RegisterUser godoc
// @Summary Register new user
// @Description Register new user
// @Tags User
// @Accept  json
// @Produce  json
// @Param request body user_payload.RegisterUserInput true "Request Body"
// @Success 200 {object} user_response.RegisterUserSwagger
// @Failure 400 {object} user_response.BaseSwagger
// @Failure 409 {object} user_response.BaseSwagger
// @Failure 500 {object} user_response.BaseSwagger
// @Router /user/register [post]
func (h *UserHandler) RegisterUser(c echo.Context) error {
	var req user_payload.RegisterUserInput
	c.Bind(&req)

	ctx, requestID := helper.GenerateRequestID()
	c.Response().Header().Set("X-Request-ID", requestID)

	id, err := h.user.RegisterUser(ctx, req.ToPB())
	if err != nil {
		return utils.ErrorResponse(c, err, map[string]interface{}{})
	}

	data := user_response.RegisterUserProtoToJsonResponse(id)

	return utils.SuccessResponse(c, constant.SuccessAddData, data)
}
