package user_response

import (
	userpb "gateway-grpc/pb/user"
)

type RegisterUserResponse struct {
	Id string `json:"id"`
}

func RegisterUserProtoToJsonResponse(in *userpb.RegisterUserResponse) RegisterUserResponse {
	return RegisterUserResponse{
		Id: in.Id,
	}
}
