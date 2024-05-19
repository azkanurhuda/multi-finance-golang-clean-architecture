package asset

import (
	"context"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/model"
)

type UseCase interface {
	Create(ctx context.Context, request *model.CreateAssetRequest) (*model.AssetResponse, error)
	GetByID(ctx context.Context, id string) (*model.AssetResponse, error)
	List(ctx context.Context) ([]model.AssetResponse, error)
	Delete(ctx context.Context, id string) (bool, error)
}
