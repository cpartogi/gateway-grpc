package user_payload

import (
	userpb "gateway-grpc/pb/user"
)

type RegisterUserInput struct {
	Email       string `json:"email"`
	FullName    string `json:"fullName"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phoneNumber"`
}

func (p *RegisterUserInput) ToPB() *userpb.RegisterUserRequest {
	return &userpb.RegisterUserRequest{
		FullName:    p.FullName,
		Email:       p.Email,
		Password:    p.Password,
		PhoneNumber: p.PhoneNumber,
	}
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (p *LoginInput) ToPB() *userpb.LoginRequest {
	return &userpb.LoginRequest{
		Email:    p.Email,
		Password: p.Password,
	}
}

type GetTokenInput struct {
	RefreahToken string `json:"refreshToken"`
}

func (p *GetTokenInput) ToPB() *userpb.GetTokenRequest {
	return &userpb.GetTokenRequest{
		RefreshToken: p.RefreahToken,
	}

}

type UpdateUserInput struct {
	UserId      string `json:"userId"`
	FullName    string `json:"fullName"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phoneNumber"`
}

func (p *UpdateUserInput) ToPB() *userpb.UpdateUserRequest {
	return &userpb.UpdateUserRequest{
		UserId:      p.UserId,
		FullName:    p.FullName,
		Password:    p.Password,
		PhoneNumber: p.PhoneNumber,
	}
}
