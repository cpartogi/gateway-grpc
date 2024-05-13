package user_response

import (
	userpb "gateway-grpc/pb/user"
	schema "gateway-grpc/schema"
)

type RegisterUserResponse struct {
	Id string `json:"id"`
}

func RegisterUserProtoToJsonResponse(in *userpb.RegisterUserResponse) RegisterUserResponse {
	return RegisterUserResponse{
		Id: in.Id,
	}
}

type RegisterUserSwagger struct {
	schema.Base
	Data RegisterUserResponse `json:"data"`
}

type BaseSwagger struct {
	schema.Base
}
