package transaction

import (
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/entity"
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	db *gorm.DB
}

func (c Transaction) UpdateCreditID(db *gorm.DB, data *entity.Transaction) error {
	return db.Model(data).Where("id = ?", data.ID).Updates(map[string]interface{}{
		"credit_id":  data.CreditID,
		"updated_at": time.Now().Local(),
	}).Error
}

func (c Transaction) FindByID(db *gorm.DB, id string) (*entity.Transaction, error) {
	var data entity.Transaction
	err := db.Preload("User").
		Preload("Asset").
		Preload("Asset.Merchant").
		Preload("Credit").
		Where("id = ?", id).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (c Transaction) Create(db *gorm.DB, transaction *entity.Transaction) error {
	return db.Create(transaction).Error
}

func (c Transaction) GetByID(db *gorm.DB, id string) (*entity.Transaction, error) {
	var data entity.Transaction
	err := db.Where("id = ?", id).First(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (c Transaction) GetList(db *gorm.DB) ([]entity.Transaction, error) {
	var data []entity.Transaction
	if err := db.Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (c Transaction) DeleteByID(db *gorm.DB, id string) error {
	result := db.Delete(&entity.Transaction{}, "id = ?", id)
	return result.Error
}

func NewTransaction(db *gorm.DB) *Transaction {
	return &Transaction{
		db: db,
	}
}
