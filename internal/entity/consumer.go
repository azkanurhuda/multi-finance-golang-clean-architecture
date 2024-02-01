package entity

import "time"

type Consumer struct {
	ID           string    `gorm:"column:id;primaryKey"`
	NIK          string    `gorm:"column:nik"`
	FullName     string    `gorm:"column:full_name"`
	LegalName    string    `gorm:"column:legal_name"`
	PlaceOfBirth string    `gorm:"column:place_of_birth"`
	DateOfBirth  time.Time `gorm:"column:date_of_birth"`
	Salary       int64     `gorm:"column:salary"`
	IDCardPhoto  string    `gorm:"column:id_card_photo"`
	SelfiePhoto  string    `gorm:"column:selfie_photo"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}

func (u *Consumer) TableName() string {
	return "consumers"
}
