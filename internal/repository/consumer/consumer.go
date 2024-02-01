package consumer

import (
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/entity"
	"gorm.io/gorm"
)

type Consumer struct {
	db *gorm.DB
}

func (c Consumer) Create(db *gorm.DB, consumer *entity.Consumer) error {
	return db.Create(consumer).Error
}

func (c Consumer) FindByNIK(db *gorm.DB, nik string) (*entity.Consumer, error) {
	var consumer entity.Consumer
	err := db.Where("nik = ?", nik).First(&consumer).Error
	if err != nil {
		return nil, err
	}

	return &consumer, nil
}

func NewConsumer(db *gorm.DB) *Consumer {
	return &Consumer{
		db: db,
	}
}
