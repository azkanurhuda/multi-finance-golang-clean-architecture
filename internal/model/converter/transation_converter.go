package converter

import (
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/entity"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/model"
	"time"
)

func TransactionToResponse(transaction *entity.Transaction) *model.TransactionResponse {
	return &model.TransactionResponse{
		ID:                transaction.ID,
		NIK:               transaction.NIK,
		ContractNumber:    transaction.ContractNumber,
		OTR:               transaction.OTR,
		AdminFee:          transaction.AdminFee,
		AmountOfInterest:  transaction.AmountOfInterest,
		InstallmentAmount: transaction.InstallmentAmount,
		AssetName:         transaction.AssetName,
		CreatedAt:         time.Now().Local(),
		UpdatedAt:         time.Now().Local(),
	}
}
