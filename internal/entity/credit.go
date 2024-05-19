package entity

import "time"

type Credit struct {
	ID                 string      `gorm:"column:id;primaryKey"`
	TransactionID      string      `gorm:"column:transaction_id"`
	Tenor              int         `gorm:"column:tenor"`
	CreditLimit        float64     `gorm:"column:credit_limit"`
	MonthlyInstallment float64     `gorm:"column:monthly_installment"`
	InterestAmount     float64     `gorm:"column:interest_amount"`
	CreatedAt          time.Time   `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt          time.Time   `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
	Transaction        Transaction `gorm:"foreignKey:transaction_id;references:id"`
}

func (c *Credit) TableName() string {
	return "credits"
}
