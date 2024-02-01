package consumer_limit

import (
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/entity"
	"gorm.io/gorm"
)

type ConsumerLimit struct {
	db *gorm.DB
}

func (c ConsumerLimit) FindByNIKAndTenor(db *gorm.DB, nik, tenor string) (*entity.ConsumerLimit, error) {
	var consumerLimit entity.ConsumerLimit
	err := db.Where("nik = ? AND tenor = ?", nik, tenor).First(&consumerLimit).Error
	if err != nil {
		return nil, err
	}

	return &consumerLimit, nil
}

func (c ConsumerLimit) Create(db *gorm.DB, consumerLimit *entity.ConsumerLimit) error {
	return db.Create(consumerLimit).Error
}

func NewConsumerLimit(db *gorm.DB) *ConsumerLimit {
	return &ConsumerLimit{
		db: db,
	}
}
