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
	Data BaseDataSwagger `json:"data"`
}

type LoginResponse struct {
	Id                    string `json:"id"`
	Token                 string `json:"token"`
	TokenExpiredAt        string `json:"tokenExpiredAt"`
	RefreshToken          string `json:"refreshToken"`
	RefreshTokenExpiredAt string `json:"refreshTokenExpiredAt,omitempty"`
}

func LoginProtoToJsonResponse(in *userpb.LoginResponse) LoginResponse {
	return LoginResponse{
		Id:                    in.Id,
		Token:                 in.Token,
		TokenExpiredAt:        in.TokenExpiredAt,
		RefreshToken:          in.RefreshToken,
		RefreshTokenExpiredAt: in.RefreshTokenExpiredAt,
	}
}

type LoginSwagger struct {
	schema.Base
	Data LoginResponse `json:"data"`
}

type BaseDataSwagger struct {
}
