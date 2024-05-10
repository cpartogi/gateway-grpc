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
