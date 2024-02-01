package model

import "time"

type UserResponse struct {
	ID        string    `json:"id,omitempty"`
	NIK       string    `json:"nik,omitempty"`
	Token     string    `json:"token,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type VerifyUserRequest struct {
	Token string `validate:"required,max=100"`
}

type RegisterUserRequest struct {
	NIK      string `json:"nik" validate:"required,max=16"`
	Password string `json:"password" validate:"required,max=100"`
}

type LoginUserRequest struct {
	NIK      string `json:"nik" validate:"required,max=16"`
	Password string `json:"password" validate:"required,max=100"`
}

type LogoutUserRequest struct {
	NIK string `json:"nik" validate:"required,max=16"`
}

type GetUserRequest struct {
	NIK string `json:"nik" validate:"required,max=16"`
}
