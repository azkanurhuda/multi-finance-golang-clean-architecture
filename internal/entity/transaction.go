package entity

import "time"

type Transaction struct {
	ID                string    `gorm:"column:id;primaryKey"`
	ContractNumber    string    `gorm:"column:contract_number"`
	NIK               string    `gorm:"column:nik"`
	OTR               int64     `gorm:"column:otr"`
	AdminFee          int64     `gorm:"column:admin_fee"`
	InstallmentAmount int64     `gorm:"column:installment_amount"`
	AmountOfInterest  int64     `gorm:"column:amount_of_interest"`
	AssetName         string    `gorm:"column:asset_name"`
	CreatedAt         time.Time `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt         time.Time `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}

func (u *Transaction) TableName() string {
	return "transactions"
}
