package credit

import (
	"context"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/model"
)

type UseCase interface {
	Create(ctx context.Context, request *model.CreateCreditRequest, auth model.Auth) (*model.CreditResponse, error)
	GetByID(ctx context.Context, id string) (*model.CreditResponse, error)
	List(ctx context.Context) ([]model.CreditResponse, error)
	Delete(ctx context.Context, id string) (bool, error)
}
