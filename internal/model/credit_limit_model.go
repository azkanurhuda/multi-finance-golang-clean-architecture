package model

import "time"

type CreditLimitResponse struct {
	ID          string        `json:"id,omitempty"`
	User        *UserResponse `json:"user,omitempty"`
	Tenor       int           `json:"tenor,omitempty"`
	CreditLimit float64       `json:"credit_limit,omitempty"`
	CreatedAt   time.Time     `json:"created_at,omitempty"`
	UpdatedAt   time.Time     `json:"updated_at,omitempty"`
}

type CreateCreditLimitRequest struct {
	UserID string  `json:"user_id" validate:"required"`
	Tenor  int     `json:"tenor" validate:"required"`
	Limit  float64 `json:"limit" validate:"required"`
}
