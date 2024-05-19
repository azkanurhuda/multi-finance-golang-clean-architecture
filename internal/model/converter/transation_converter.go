package converter

import (
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/entity"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/model"
)

func TransactionToResponse(transaction *entity.Transaction) *model.TransactionResponse {
	var creditResponse *model.CreditResponse
	if transaction.Credit != nil {
		creditResponse = &model.CreditResponse{
			ID:                 transaction.Credit.ID,
			Tenor:              transaction.Credit.Tenor,
			CreditLimit:        transaction.Credit.CreditLimit,
			MonthlyInstallment: transaction.Credit.MonthlyInstallment,
			InterestAmount:     transaction.Credit.InterestAmount,
			CreatedAt:          transaction.Credit.CreatedAt,
			UpdatedAt:          transaction.Credit.UpdatedAt,
		}

		return &model.TransactionResponse{
			ID:             transaction.ID,
			User:           UserToResponse(&transaction.User),
			Asset:          AssetToResponse(&transaction.Asset, transaction.Asset.Merchant),
			ContractNumber: transaction.ContractNumber,
			TotalPayment:   transaction.TotalPayment,
			PaymentMethod:  transaction.PaymentMethod,
			CreatedAt:      transaction.CreatedAt,
			UpdatedAt:      transaction.UpdatedAt,
			Credit:         creditResponse,
		}
	}

	return &model.TransactionResponse{
		ID:             transaction.ID,
		User:           UserToResponse(&transaction.User),
		Asset:          AssetToResponse(&transaction.Asset, transaction.Asset.Merchant),
		ContractNumber: transaction.ContractNumber,
		TotalPayment:   transaction.TotalPayment,
		PaymentMethod:  transaction.PaymentMethod,
		CreatedAt:      transaction.CreatedAt,
		UpdatedAt:      transaction.UpdatedAt,
	}
}
