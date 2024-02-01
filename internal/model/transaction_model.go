package model

import "time"

type CreateTransactionRequest struct {
	OTR               int64  `json:"otr" validate:"required"`
	InstallmentAmount int64  `json:"installment_amount" validate:"required"`
	AssetName         string `json:"asset_name" validate:"required"`
}

type TransactionResponse struct {
	ID                string    `json:"id,omitempty"`
	NIK               string    `json:"nik,omitempty"`
	ContractNumber    string    `json:"contract_number,omitempty"`
	OTR               int64     `json:"otr,omitempty"`
	AdminFee          int64     `json:"admin_fee,omitempty"`
	AmountOfInterest  int64     `json:"amount_of_interest,omitempty"`
	InstallmentAmount int64     `json:"installment_amount"`
	AssetName         string    `json:"asset_name,omitempty"`
	CreatedAt         time.Time `json:"created_at,omitempty"`
	UpdatedAt         time.Time `json:"updated_at,omitempty"`
}
