package repository

import (
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/entity"
	"gorm.io/gorm"
)

type User interface {
	FindByToken(db *gorm.DB, token string) (*entity.User, error)
	FindByNIK(db *gorm.DB, nik string) (*entity.User, error)
	Create(db *gorm.DB, user *entity.User) error
	UpdateTokenByNIK(db *gorm.DB, user *entity.User, nik string) error
}

type Consumer interface {
	Create(db *gorm.DB, consumer *entity.Consumer) error
	FindByNIK(db *gorm.DB, nik string) (*entity.Consumer, error)
}

type ConsumerLimit interface {
	Create(db *gorm.DB, consumerLimit *entity.ConsumerLimit) error
	FindByNIKAndTenor(db *gorm.DB, nik, tenor string) (*entity.ConsumerLimit, error)
}

type Transaction interface {
	Create(db *gorm.DB, consumerLimit *entity.Transaction) error
}
