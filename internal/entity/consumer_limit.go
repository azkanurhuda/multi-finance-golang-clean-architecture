package entity

import "time"

type ConsumerLimit struct {
	ID        string    `gorm:"column:id;primaryKey"`
	NIK       string    `gorm:"column:nik"`
	Tenor     string    `gorm:"column:tenor"`
	Limits    int64     `gorm:"column:limits"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}

func (u *ConsumerLimit) TableName() string {
	return "consumer_limits"
}
