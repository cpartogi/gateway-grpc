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

type UserResponse struct {
	Id          string `json:"id"`
	Email       string `json:"email"`
	FullName    string `json:"fullName"`
	PhoneNumber string `json:"phoneNumber"`
}

func UserProtoToJsonResponse(in *userpb.UserResponse) UserResponse {
	return UserResponse{
		Id:          in.Id,
		Email:       in.Email,
		FullName:    in.FullName,
		PhoneNumber: in.PhoneNumber,
	}
}

type UserSwagger struct {
	schema.Base
	Data UserResponse `json:"data"`
}

type UpdateUserResponse struct {
	Id string `json:"id"`
}

func UpdateUserProtoToJsonResponse(in *userpb.UpdateUserResponse) UpdateUserResponse {
	return UpdateUserResponse{
		Id: in.Id,
	}
}

type UpdateUserSwagger struct {
	schema.Base
	Data UpdateUserResponse `json:"data"`
}
