package credit_limit

import (
	"context"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/model"
)

type UseCase interface {
	Create(ctx context.Context, request *model.CreateCreditLimitRequest) (*model.CreditLimitResponse, error)
	GetByID(ctx context.Context, id string) (*model.CreditLimitResponse, error)
	List(ctx context.Context) ([]model.CreditLimitResponse, error)
	Delete(ctx context.Context, id string) (bool, error)
}
