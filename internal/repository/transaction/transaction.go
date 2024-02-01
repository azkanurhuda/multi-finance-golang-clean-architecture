package transaction

import (
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/entity"
	"gorm.io/gorm"
)

type Transaction struct {
	db *gorm.DB
}

func (c Transaction) Create(db *gorm.DB, transaction *entity.Transaction) error {
	return db.Create(transaction).Error
}

func NewTransaction(db *gorm.DB) *Transaction {
	return &Transaction{
		db: db,
	}
}
