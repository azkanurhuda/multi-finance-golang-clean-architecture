package merchant

import (
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/entity"
	"gorm.io/gorm"
)

type Merchant struct {
	db *gorm.DB
}

func (m Merchant) Create(db *gorm.DB, data *entity.Merchant) error {
	return db.Create(data).Error
}

func (m Merchant) GetByID(db *gorm.DB, id string) (*entity.Merchant, error) {
	var data entity.Merchant
	err := db.Where("id = ?", id).First(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (m Merchant) GetList(db *gorm.DB) ([]entity.Merchant, error) {
	var data []entity.Merchant
	if err := db.Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (m Merchant) DeleteByID(db *gorm.DB, id string) error {
	result := db.Delete(&entity.Merchant{}, "id = ?", id)
	return result.Error
}

func NewMerchant(db *gorm.DB) *Merchant {
	return &Merchant{
		db: db,
	}
}
