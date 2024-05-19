package model

import "time"

type AssetResponse struct {
	ID          string            `json:"id,omitempty"`
	Merchant    *MerchantResponse `json:"merchant,omitempty"`
	Name        string            `json:"name,omitempty"`
	OTR         float64           `json:"otr,omitempty"`
	AdminFee    float64           `json:"admin_fee,omitempty"`
	Description string            `json:"description,omitempty"`
	CreatedAt   time.Time         `json:"created_at,omitempty"`
	UpdatedAt   time.Time         `json:"updated_at,omitempty"`
}

type CreateAssetRequest struct {
	MerchantID  string  `json:"merchant_id" validate:"required"`
	Name        string  `json:"name" validate:"required"`
	OTR         float64 `json:"otr" validate:"required"`
	AdminFee    float64 `json:"admin_fee" validate:"required"`
	Description string  `json:"description" validate:"required"`
}
