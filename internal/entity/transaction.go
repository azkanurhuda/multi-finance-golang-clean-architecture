package entity

import "time"

type Transaction struct {
	ID             string    `gorm:"column:id;primaryKey"`
	UserID         string    `gorm:"column:user_id"`
	AssetID        string    `gorm:"column:asset_id"`
	ContractNumber string    `gorm:"column:contract_number"`
	TotalPayment   float64   `gorm:"column:total_payment"`
	PaymentMethod  string    `gorm:"column:payment_method"`
	CreatedAt      time.Time `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt      time.Time `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
	CreditID       *string   `gorm:"column:credit_id"`
	User           User      `gorm:"foreignKey:user_id;references:id"`
	Asset          Asset     `gorm:"foreignKey:asset_id;references:id"`
	Credit         *Credit   `gorm:"foreignKey:credit_id;references:id"`
}

func (t *Transaction) TableName() string {
	return "transactions"
}
