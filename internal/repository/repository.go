package repository

import (
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/repository/asset"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/repository/consumer"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/repository/credit"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/repository/credit_limit"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/repository/credit_payment"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/repository/merchant"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/repository/transaction"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/repository/user"
	"gorm.io/gorm"
)

type Repository struct {
	db            *gorm.DB
	User          User
	Consumer      Consumer
	Merchant      Merchant
	Asset         Asset
	Credit        Credit
	CreditLimit   CreditLimit
	Transaction   Transaction
	CreditPayment CreditPayment
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db:            db,
		User:          user.NewUser(db),
		Consumer:      consumer.NewConsumer(db),
		Merchant:      merchant.NewMerchant(db),
		Asset:         asset.NewAsset(db),
		Credit:        credit.NewCredit(db),
		CreditLimit:   credit_limit.NewCreditLimit(db),
		Transaction:   transaction.NewTransaction(db),
		CreditPayment: credit_payment.NewCreditPayment(db),
	}
}
