package consumer

import (
	"context"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/model"
)

type UseCase interface {
	Create(ctx context.Context, request *model.CreateConsumerRequest, auth model.Auth) (*model.ConsumerResponse, error)
	CreateByAdmin(ctx context.Context, request *model.CreateConsumerByAdminRequest) (*model.ConsumerResponse, error)
}
