package converter

import (
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/entity"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/model"
)

func CreditToResponse(credit *entity.Credit) *model.CreditResponse {
	return &model.CreditResponse{
		ID:                 credit.ID,
		Transaction:        TransactionToResponse(&credit.Transaction),
		Tenor:              credit.Tenor,
		CreditLimit:        credit.CreditLimit,
		MonthlyInstallment: credit.MonthlyInstallment,
		InterestAmount:     credit.InterestAmount,
		CreatedAt:          credit.CreatedAt,
		UpdatedAt:          credit.UpdatedAt,
	}
}

func ListCreditToResponse(credit []entity.Credit) []model.CreditResponse {
	var data []model.CreditResponse

	for _, v := range credit {
		data = append(data, model.CreditResponse{
			ID:                 v.ID,
			Transaction:        nil,
			Tenor:              v.Tenor,
			CreditLimit:        v.CreditLimit,
			MonthlyInstallment: v.MonthlyInstallment,
			InterestAmount:     v.InterestAmount,
			CreatedAt:          v.CreatedAt,
			UpdatedAt:          v.UpdatedAt,
		})
	}

	return data
}
