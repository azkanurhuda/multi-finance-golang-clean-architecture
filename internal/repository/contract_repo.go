package repository

import (
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/entity"
	"gorm.io/gorm"
)

type User interface {
	FindByToken(db *gorm.DB, token string) (*entity.User, error)
	FindByID(db *gorm.DB, id string) (*entity.User, error)
	FindByEmail(db *gorm.DB, email string) (*entity.User, error)
	Create(db *gorm.DB, user *entity.User) error
	UpdateTokenByEmail(db *gorm.DB, user *entity.User, email string) error
}

type Consumer interface {
	Create(db *gorm.DB, consumer *entity.Consumer) error
	FindByUserID(db *gorm.DB, userID string) (*entity.Consumer, error)
	GetByID(db *gorm.DB, id string) (*entity.Consumer, error)
	GetList(db *gorm.DB) ([]entity.Consumer, error)
	DeleteByID(db *gorm.DB) error
	FindByImage(db *gorm.DB, target, image string) (*entity.Consumer, error)
}

type Merchant interface {
	Create(db *gorm.DB, data *entity.Merchant) error
	GetByID(db *gorm.DB, id string) (*entity.Merchant, error)
	GetList(db *gorm.DB) ([]entity.Merchant, error)
	DeleteByID(db *gorm.DB, id string) error
}

type Asset interface {
	Create(db *gorm.DB, data *entity.Asset) error
	GetByID(db *gorm.DB, id string) (*entity.Asset, error)
	GetList(db *gorm.DB) ([]entity.Asset, error)
	DeleteByID(db *gorm.DB, id string) error
}

type Credit interface {
	Create(db *gorm.DB, data *entity.Credit) error
	GetByID(db *gorm.DB, id string) (*entity.Credit, error)
	GetList(db *gorm.DB) ([]entity.Credit, error)
	DeleteByID(db *gorm.DB, id string) error
}

type CreditLimit interface {
	Create(db *gorm.DB, data *entity.CreditLimit) error
	GetByID(db *gorm.DB, id string) (*entity.CreditLimit, error)
	GetList(db *gorm.DB) ([]entity.CreditLimit, error)
	DeleteByID(db *gorm.DB, id string) error
	GetByUserIDAndTenor(db *gorm.DB, userID string, tenor int) (*entity.CreditLimit, error)
}

type Transaction interface {
	Create(db *gorm.DB, data *entity.Transaction) error
	GetByID(db *gorm.DB, id string) (*entity.Transaction, error)
	GetList(db *gorm.DB) ([]entity.Transaction, error)
	DeleteByID(db *gorm.DB, id string) error
	FindByID(db *gorm.DB, id string) (*entity.Transaction, error)
	UpdateCreditID(db *gorm.DB, data *entity.Transaction) error
}

type CreditPayment interface {
	Create(db *gorm.DB, data *entity.CreditPayment) error
	GetByID(db *gorm.DB, id string) (*entity.CreditPayment, error)
	GetList(db *gorm.DB) ([]entity.CreditPayment, error)
	DeleteByID(db *gorm.DB, id string) error
	GetListByCreditID(db *gorm.DB, id string) ([]entity.CreditPayment, error)
}
