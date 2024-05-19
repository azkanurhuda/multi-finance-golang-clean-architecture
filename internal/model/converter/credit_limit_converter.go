package converter

import (
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/entity"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/model"
)

func CreditLimitToResponse(creditLimit *entity.CreditLimit, user *entity.User) *model.CreditLimitResponse {
	return &model.CreditLimitResponse{
		ID:          creditLimit.ID,
		User:        UserToResponse(user),
		Tenor:       creditLimit.Tenor,
		CreditLimit: creditLimit.CreditLimit,
		CreatedAt:   creditLimit.CreatedAt,
		UpdatedAt:   creditLimit.UpdatedAt,
	}
}

func ListCreditLimitToResponse(credit []entity.CreditLimit) []model.CreditLimitResponse {
	var data []model.CreditLimitResponse

	for _, v := range credit {
		data = append(data, model.CreditLimitResponse{
			ID:          v.ID,
			User:        nil,
			Tenor:       v.Tenor,
			CreditLimit: v.CreditLimit,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
		})
	}

	return data
}
