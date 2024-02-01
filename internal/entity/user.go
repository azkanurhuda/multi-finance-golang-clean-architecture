package entity

import "time"

type User struct {
	ID        string    `gorm:"column:id;primaryKey"`
	NIK       string    `gorm:"column:nik"`
	Password  string    `gorm:"column:password"`
	Token     string    `gorm:"column:token"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}

func (u *User) TableName() string {
	return "users"
}
