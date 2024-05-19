package model

import "time"

type CreditPaymentResponse struct {
	ID            string          `json:"id,omitempty"`
	Credit        *CreditResponse `json:"credit,omitempty"`
	PaymentAmount float64         `json:"payment_amount,omitempty"`
	CreatedAt     time.Time       `json:"created_at,omitempty"`
	UpdatedAt     time.Time       `json:"updated_at,omitempty"`
}

type CreateCreditPaymentRequest struct {
	CreditID      string  `json:"credit_id" validate:"required"`
	PaymentAmount float64 `json:"payment_amount" validate:"required"`
}
