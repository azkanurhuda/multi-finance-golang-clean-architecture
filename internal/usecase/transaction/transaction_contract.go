package transaction

import (
	"context"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/model"
)

type UseCase interface {
	Create(ctx context.Context, request *model.CreateTransactionRequest, auth model.Auth) (*model.TransactionResponse, error)
	CreateByAdmin(ctx context.Context, request *model.CreateTransactionByAdminRequest) (*model.TransactionResponse, error)
}
