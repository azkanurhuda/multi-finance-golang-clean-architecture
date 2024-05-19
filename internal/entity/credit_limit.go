package entity

import "time"

type CreditLimit struct {
	ID          string    `gorm:"column:id;primaryKey"`
	UserID      string    `gorm:"column:user_id"`
	Tenor       int       `gorm:"column:tenor"`
	CreditLimit float64   `gorm:"column:credit_limit"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
	User        User      `gorm:"foreignKey:user_id;references:id"`
}

func (c *CreditLimit) TableName() string {
	return "credit_limits"
}
