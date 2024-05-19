package asset

import (
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/entity"
	"gorm.io/gorm"
)

type Asset struct {
	db *gorm.DB
}

func (a Asset) Create(db *gorm.DB, data *entity.Asset) error {
	return db.Create(data).Error
}

func (a Asset) GetByID(db *gorm.DB, id string) (*entity.Asset, error) {
	var data entity.Asset
	err := db.Where("id = ?", id).First(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (a Asset) GetList(db *gorm.DB) ([]entity.Asset, error) {
	var data []entity.Asset
	if err := db.Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (a Asset) DeleteByID(db *gorm.DB, id string) error {
	result := db.Delete(&entity.Asset{}, "id = ?", id)
	return result.Error
}

func NewAsset(db *gorm.DB) *Asset {
	return &Asset{
		db: db,
	}
}
