package credit

import (
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/entity"
	"gorm.io/gorm"
)

type Credit struct {
	db *gorm.DB
}

func (c Credit) Create(db *gorm.DB, data *entity.Credit) error {
	return db.Create(data).Error
}

func (c Credit) GetByID(db *gorm.DB, id string) (*entity.Credit, error) {
	var data entity.Credit
	err := db.
		Preload("Transaction").
		Preload("Transaction.User").
		Preload("Transaction.Asset").
		Preload("Transaction.Asset.Merchant").
		Where("id = ?", id).First(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (c Credit) GetList(db *gorm.DB) ([]entity.Credit, error) {
	var data []entity.Credit
	if err := db.Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (c Credit) DeleteByID(db *gorm.DB, id string) error {
	result := db.Delete(&entity.Credit{}, "id = ?", id)
	return result.Error
}

func NewCredit(db *gorm.DB) *Credit {
	return &Credit{
		db: db,
	}
}
