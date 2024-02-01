package repository

import (
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/repository/consumer"
	consumer_limit "github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/repository/consumer_limit"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/repository/transaction"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/repository/user"
	"gorm.io/gorm"
)

type Repository struct {
	db            *gorm.DB
	User          User
	Consumer      Consumer
	ConsumerLimit ConsumerLimit
	Transaction   Transaction
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db:            db,
		User:          user.NewUser(db),
		Consumer:      consumer.NewConsumer(db),
		ConsumerLimit: consumer_limit.NewConsumerLimit(db),
		Transaction:   transaction.NewTransaction(db),
	}
}
