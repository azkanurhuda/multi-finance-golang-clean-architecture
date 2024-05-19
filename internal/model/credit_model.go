package model

import "time"

type CreditResponse struct {
	ID                 string               `json:"id,omitempty"`
	Transaction        *TransactionResponse `json:"transaction,omitempty"`
	Tenor              int                  `json:"tenor,omitempty"`
	CreditLimit        float64              `json:"credit_limit,omitempty"`
	MonthlyInstallment float64              `json:"monthly_installment,omitempty"`
	InterestAmount     float64              `json:"interest_amount,omitempty"`
	CreatedAt          time.Time            `json:"created_at,omitempty"`
	UpdatedAt          time.Time            `json:"updated_at,omitempty"`
}

type CreateCreditRequest struct {
	TransactionID string `json:"transaction_id" validate:"required"`
	Tenor         int    `json:"tenor" validate:"required"`
}
