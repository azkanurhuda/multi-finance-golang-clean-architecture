package model

import "time"

type CreateConsumerRequest struct {
	FullName     string `json:"full_name" validate:"required,max=255"`
	LegalName    string `json:"legal_name" validate:"required,max=255"`
	PlaceOfBirth string `json:"place_of_birth"`
	DateOfBirth  string `json:"date_of_birth"`
	Salary       int64  `json:"salary"`
	IDCardPhoto  string `json:"id_card_photo"`
	SelfiePhoto  string `json:"selfie_photo"`
}

type ConsumerResponse struct {
	ID           string    `json:"id,omitempty"`
	NIK          string    `json:"nik,omitempty"`
	FullName     string    `json:"full_name,omitempty"`
	LegalName    string    `json:"legal_name,omitempty"`
	PlaceOfBirth string    `json:"place_of_birth,omitempty"`
	DateOfBirth  time.Time `json:"date_of_birth,omitempty"`
	Salary       int64     `json:"salary,omitempty"`
	IDCardPhoto  string    `json:"id_card_photo,omitempty"`
	SelfiePhoto  string    `json:"selfie_photo,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
}
