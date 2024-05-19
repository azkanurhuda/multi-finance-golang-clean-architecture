package entity

import "time"

type User struct {
	ID             string     `gorm:"column:id;primaryKey"`
	Email          string     `gorm:"column:email"`
	Password       string     `gorm:"column:password"`
	Role           string     `gorm:"column:role"`
	Token          string     `gorm:"column:token"`
	TokenExpiredAt *time.Time `gorm:"column:token_expired_at"`
	CreatedAt      time.Time  `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt      time.Time  `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}

func (u *User) TableName() string {
	return "users"
}
