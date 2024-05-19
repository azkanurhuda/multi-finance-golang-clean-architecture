package model

import "time"

type MerchantResponse struct {
	ID        string    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type CreateMerchantRequest struct {
	Name string `json:"name" validate:"required"`
}
