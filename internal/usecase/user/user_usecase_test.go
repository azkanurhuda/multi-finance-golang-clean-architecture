package user

import (
	"context"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/model"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"reflect"
	"testing"
)

func TestUserUseCase_Create(t *testing.T) {
	type fields struct {
		DB         *gorm.DB
		Log        *logrus.Logger
		Validate   *validator.Validate
		Repository *repository.Repository
	}
	type args struct {
		ctx     context.Context
		request *model.RegisterUserRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.UserResponse
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := UserUseCase{
				DB:         tt.fields.DB,
				Log:        tt.fields.Log,
				Validate:   tt.fields.Validate,
				Repository: tt.fields.Repository,
			}
			got, err := c.Create(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}
