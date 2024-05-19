package merchant

import (
	"context"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/model"
)

type UseCase interface {
	Create(ctx context.Context, request *model.CreateMerchantRequest) (*model.MerchantResponse, error)
	GetByID(ctx context.Context, id string) (*model.MerchantResponse, error)
	List(ctx context.Context) ([]model.MerchantResponse, error)
	Delete(ctx context.Context, id string) (bool, error)
}
