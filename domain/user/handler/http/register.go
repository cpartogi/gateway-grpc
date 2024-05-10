package http

import (
	"fmt"
	user_payload "gateway-grpc/domain/user/payload"
	"gateway-grpc/lib/constant"
	"gateway-grpc/lib/pkg/utils"

	user_response "gateway-grpc/domain/user/response"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/status"
)

func (h *UserHandler) RegisterUser(c echo.Context) error {
	ctx := c.Request().Context()
	var req user_payload.RegisterUserInput

	c.Bind(&req)

	id, err := h.user.RegisterUser(ctx, req.ToPB())
	if err != nil {
		st, _ := status.FromError(err)
		fmt.Println(st.Code())
		return utils.ErrorResponse(c, err, map[string]interface{}{})
	}

	data := user_response.RegisterUserProtoToJsonResponse(id)

	return utils.SuccessResponse(c, constant.SuccessAddData, data)
}
