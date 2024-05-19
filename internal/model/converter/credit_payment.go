package converter

import (
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/entity"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/model"
)

func CreditPaymentToResponse(creditPayment *entity.CreditPayment) *model.CreditPaymentResponse {
	return &model.CreditPaymentResponse{
		ID:            creditPayment.ID,
		Credit:        CreditToResponse(&creditPayment.Credit),
		PaymentAmount: creditPayment.PaymentAmount,
		CreatedAt:     creditPayment.CreatedAt,
		UpdatedAt:     creditPayment.UpdatedAt,
	}
}

func ListCreditPaymentToResponse(creditPayment []entity.CreditPayment) []model.CreditPaymentResponse {
	var data []model.CreditPaymentResponse
	for _, v := range creditPayment {
		data = append(data, model.CreditPaymentResponse{
			ID:            v.CreditID,
			Credit:        nil,
			PaymentAmount: v.PaymentAmount,
			CreatedAt:     v.CreatedAt,
			UpdatedAt:     v.UpdatedAt,
		})
	}
	return data
}
