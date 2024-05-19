package credit_payment

import (
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/entity"
	"gorm.io/gorm"
)

type CreditPayment struct {
	db *gorm.DB
}

func (c CreditPayment) GetListByCreditID(db *gorm.DB, id string) ([]entity.CreditPayment, error) {
	var data []entity.CreditPayment
	err := db.
		Preload("Credit").
		Preload("Credit.Transaction").
		Preload("Credit.Transaction.User").
		Preload("Credit.Transaction.Asset").
		Preload("Credit.Transaction.Asset.Merchant").
		Where("credit_id = ?", id).Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c CreditPayment) Create(db *gorm.DB, data *entity.CreditPayment) error {
	return db.Create(data).Error
}

func (c CreditPayment) GetByID(db *gorm.DB, id string) (*entity.CreditPayment, error) {
	var data entity.CreditPayment
	err := db.
		Preload("Credit").
		Preload("Credit.Transaction").
		Preload("Credit.Transaction.User").
		Preload("Credit.Transaction.Asset").
		Preload("Credit.Transaction.Asset.Merchant").
		Where("id = ?", id).First(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (c CreditPayment) GetList(db *gorm.DB) ([]entity.CreditPayment, error) {
	var data []entity.CreditPayment
	if err := db.Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (c CreditPayment) DeleteByID(db *gorm.DB, id string) error {
	result := db.Delete(&entity.CreditPayment{}, "id = ?", id)
	return result.Error
}

func NewCreditPayment(db *gorm.DB) *CreditPayment {
	return &CreditPayment{
		db: db,
	}
}
