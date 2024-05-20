package http

import (
	user_payload "gateway-grpc/domain/user/payload"
	user_response "gateway-grpc/domain/user/response"
	"gateway-grpc/lib/constant"
	"gateway-grpc/lib/helper"
	"gateway-grpc/lib/pkg/utils"

	"github.com/labstack/echo/v4"
)

// Update User Profile godoc
// @Summary update user profile
// @Description update user profile
// @Tags User
// @Accept  json
// @Produce  json
// @Param request body user_payload.UpdateUserInput true "Request Body"
// @Success 200 {object} user_response.UpdateUserSwagger
// @Failure 400 {object} user_response.BaseSwagger
// @Failure 403 {object} user_response.BaseSwagger
// @Failure 500 {object} user_response.BaseSwagger
// @Router /user/profile [put]
// @Param  X-Authorization header string true "X-Authorization" default(JWT <token-from-login>) example(JWT eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c)
func (h *UserHandler) UpdateUser(c echo.Context) error {
	var req user_payload.UpdateUserInput
	c.Bind(&req)

	ctx, requestID := helper.GenerateRequestID(c)
	c.Response().Header().Set("X-Request-ID", requestID)

	user, err := h.user.UpdateUser(ctx, req.ToPB())
	if err != nil {
		return utils.ErrorResponse(c, err, map[string]interface{}{})
	}

	data := user_response.UpdateUserProtoToJsonResponse(user)

	return utils.SuccessResponse(c, constant.SuccessUpdateData, data)
}
