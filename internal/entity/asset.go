package entity

import "time"

type Asset struct {
	ID          string    `gorm:"column:id;primaryKey"`
	MerchantID  string    `gorm:"column:merchant_id"`
	Name        string    `gorm:"column:name"`
	OTR         float64   `gorm:"column:otr"`
	AdminFee    float64   `gorm:"column:admin_fee"`
	Description string    `gorm:"column:description"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
	Merchant    *Merchant `gorm:"foreignKey:merchant_id;references:id"`
}

func (a *Asset) TableName() string {
	return "assets"
}
