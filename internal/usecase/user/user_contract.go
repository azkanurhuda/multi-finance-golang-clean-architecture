package user

import (
	"context"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/model"
)

type UseCase interface {
	Create(ctx context.Context, request *model.RegisterUserRequest) (*model.UserResponse, error)
	Verify(ctx context.Context, request *model.VerifyUserRequest) (*model.Auth, error)
	Login(ctx context.Context, request *model.LoginUserRequest) (*model.UserResponse, error)
	Current(ctx context.Context, request *model.GetUserRequest) (*model.UserResponse, error)
	Logout(ctx context.Context, request *model.LogoutUserRequest) (bool, error)
}
