package model

import "time"

type TransactionResponse struct {
	ID             string          `json:"id,omitempty"`
	User           *UserResponse   `json:"user,omitempty"`
	Asset          *AssetResponse  `json:"asset,omitempty"`
	ContractNumber string          `json:"contract_number,omitempty"`
	TotalPayment   float64         `json:"total_payment,omitempty"`
	PaymentMethod  string          `json:"payment_method,omitempty"`
	CreatedAt      time.Time       `json:"created_at,omitempty"`
	UpdatedAt      time.Time       `json:"updated_at,omitempty"`
	Credit         *CreditResponse `json:"credit,omitempty"`
}

type CreateTransactionRequest struct {
	AssetID        string  `json:"asset_id" validate:"required"`
	ContractNumber string  `json:"contract_number" validate:"required"`
	PaymentMethod  string  `json:"payment_method" validate:"required"`
	CreditID       *string `json:"credit_id"`
}

type CreateTransactionByAdminRequest struct {
	UserID         string  `json:"user_id" validate:"required"`
	AssetID        string  `json:"asset_id" validate:"required"`
	ContractNumber string  `json:"contract_number" validate:"required"`
	PaymentMethod  string  `json:"payment_method" validate:"required"`
	CreditID       *string `json:"credit_id"`
}
