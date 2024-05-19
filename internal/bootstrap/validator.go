package bootstrap

import (
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/config"
	"github.com/go-playground/validator/v10"
)

func NewValidator(config *config.Config) *validator.Validate {
	return validator.New()
}
