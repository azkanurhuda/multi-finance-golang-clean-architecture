package entity

import "time"

type CreditPayment struct {
	ID            string    `gorm:"column:id;primaryKey"`
	CreditID      string    `gorm:"column:credit_id"`
	PaymentAmount float64   `gorm:"column:payment_amount"`
	CreatedAt     time.Time `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt     time.Time `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
	Credit        Credit    `gorm:"foreignKey:credit_id;references:id"`
}

func (c *CreditPayment) TableName() string {
	return "credit_payments"
}
