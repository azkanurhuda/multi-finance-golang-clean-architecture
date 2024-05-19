package credit_payment

import (
	"context"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/model"
)

type UseCase interface {
	Create(ctx context.Context, request *model.CreateCreditPaymentRequest) (*model.CreditPaymentResponse, error)
	GetByID(ctx context.Context, id string) (*model.CreditPaymentResponse, error)
	List(ctx context.Context) ([]model.CreditPaymentResponse, error)
	Delete(ctx context.Context, id string) (bool, error)
}
