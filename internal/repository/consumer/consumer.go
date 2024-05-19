package consumer

import (
	"fmt"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/entity"
	"gorm.io/gorm"
)

type Consumer struct {
	db *gorm.DB
}

func (c Consumer) FindByImage(db *gorm.DB, target, image string) (*entity.Consumer, error) {
	var consumer entity.Consumer
	err := db.Where(fmt.Sprintf("%s = ?", target), image).First(&consumer).Error
	if err != nil {
		return nil, err
	}

	return &consumer, nil

}

func (c Consumer) Create(db *gorm.DB, consumer *entity.Consumer) error {
	return db.Create(consumer).Error
}

func (c Consumer) FindByUserID(db *gorm.DB, userID string) (*entity.Consumer, error) {
	var consumer entity.Consumer
	err := db.Where("user_id = ?", userID).First(&consumer).Error
	if err != nil {
		return nil, err
	}

	return &consumer, nil
}

func (c Consumer) GetByID(db *gorm.DB, id string) (*entity.Consumer, error) {
	var consumer entity.Consumer
	err := db.Where("id = ?", id).First(&consumer).Error
	if err != nil {
		return nil, err
	}

	return &consumer, nil
}

func (c Consumer) GetList(db *gorm.DB) ([]entity.Consumer, error) {
	var data []entity.Consumer
	if err := db.Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (c Consumer) DeleteByID(db *gorm.DB) error {
	var data entity.Consumer
	return db.Delete(data).Error
}

func NewConsumer(db *gorm.DB) *Consumer {
	return &Consumer{
		db: db,
	}
}
