package model

import "time"

type ConsumerResponse struct {
	ID           string        `json:"id,omitempty"`
	User         *UserResponse `json:"user,omitempty"`
	NIK          string        `json:"nik,omitempty"`
	FullName     string        `json:"full_name,omitempty"`
	LegalName    string        `json:"legal_name,omitempty"`
	PhoneNumber  string        `json:"phone_number,omitempty"`
	Address      string        `json:"address,omitempty"`
	PlaceOfBirth string        `json:"place_of_birth,omitempty"`
	DateOfBirth  time.Time     `json:"date_of_birth,omitempty"`
	Salary       float64       `json:"salary,omitempty"`
	IDCardPhoto  string        `json:"id_card_photo,omitempty"`
	SelfiePhoto  string        `json:"selfie_photo,omitempty"`
	CreatedAt    time.Time     `json:"created_at,omitempty"`
	UpdatedAt    time.Time     `json:"updated_at,omitempty"`
}

type CreateConsumerByAdminRequest struct {
	NIK          string  `json:"nik" validate:"required"`
	UserID       string  `json:"user_id" validate:"required"`
	FullName     string  `json:"full_name" validate:"required,max=255"`
	LegalName    string  `json:"legal_name" validate:"required,max=255"`
	PhoneNumber  string  `json:"phone_number" validate:"required"`
	Address      string  `json:"address" validate:"required"`
	PlaceOfBirth string  `json:"place_of_birth" validate:"required,max=100"`
	DateOfBirth  string  `json:"date_of_birth" validate:"required"`
	Salary       float64 `json:"salary" validate:"required"`
	IDCardPhoto  string  `json:"id_card_photo" validate:"required"`
	SelfiePhoto  string  `json:"selfie_photo" validate:"required"`
}

type CreateConsumerRequest struct {
	NIK          string  `json:"nik" validate:"required"`
	FullName     string  `json:"full_name" validate:"required,max=255"`
	LegalName    string  `json:"legal_name" validate:"required,max=255"`
	PhoneNumber  string  `json:"phone_number" validate:"required"`
	Address      string  `json:"address" validate:"required"`
	PlaceOfBirth string  `json:"place_of_birth" validate:"required,max=100"`
	DateOfBirth  string  `json:"date_of_birth" validate:"required"`
	Salary       float64 `json:"salary" validate:"required"`
	IDCardPhoto  string  `json:"id_card_photo" validate:"required"`
	SelfiePhoto  string  `json:"selfie_photo" validate:"required"`
}
