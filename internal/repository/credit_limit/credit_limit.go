package credit_limit

import (
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/entity"
	"gorm.io/gorm"
)

type CreditLimit struct {
	db *gorm.DB
}

func (c CreditLimit) GetByUserIDAndTenor(db *gorm.DB, userID string, tenor int) (*entity.CreditLimit, error) {
	var data entity.CreditLimit
	err := db.Where("user_id = ? and tenor = ?", userID, tenor).First(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (c CreditLimit) Create(db *gorm.DB, data *entity.CreditLimit) error {
	return db.Create(data).Error
}

func (c CreditLimit) GetByID(db *gorm.DB, id string) (*entity.CreditLimit, error) {
	var data entity.CreditLimit
	err := db.Where("id = ?", id).First(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (c CreditLimit) GetList(db *gorm.DB) ([]entity.CreditLimit, error) {
	var data []entity.CreditLimit
	if err := db.Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (c CreditLimit) DeleteByID(db *gorm.DB, id string) error {
	result := db.Delete(&entity.CreditLimit{}, "id = ?", id)
	return result.Error
}

func NewCreditLimit(db *gorm.DB) *CreditLimit {
	return &CreditLimit{
		db: db,
	}
}
