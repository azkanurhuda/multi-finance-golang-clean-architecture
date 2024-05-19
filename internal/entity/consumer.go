package entity

import "time"

type Consumer struct {
	ID           string    `gorm:"column:id;primaryKey"`
	UserID       string    `gorm:"column:user_id"`
	NIK          string    `gorm:"column:nik"`
	FullName     string    `gorm:"column:full_name"`
	LegalName    string    `gorm:"column:legal_name"`
	PhoneNumber  string    `gorm:"column:phone_number"`
	Address      string    `gorm:"column:address"`
	PlaceOfBirth string    `gorm:"column:place_of_birth"`
	DateOfBirth  time.Time `gorm:"column:date_of_birth"`
	Salary       float64   `gorm:"column:salary"`
	IDCardPhoto  string    `gorm:"column:id_card_photo"`
	SelfiePhoto  string    `gorm:"column:selfie_photo"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
	User         User      `gorm:"foreignKey:user_id;references:id"`
}

func (c *Consumer) TableName() string {
	return "consumers"
}
