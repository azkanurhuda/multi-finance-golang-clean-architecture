package converter

import (
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/entity"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/model"
)

func MerchantToResponse(merchant *entity.Merchant) *model.MerchantResponse {
	return &model.MerchantResponse{
		ID:        merchant.ID,
		Name:      merchant.Name,
		CreatedAt: merchant.CreatedAt,
		UpdatedAt: merchant.UpdatedAt,
	}
}

func ListMerchantToResponse(merchant []entity.Merchant) []model.MerchantResponse {
	var data []model.MerchantResponse
	for _, v := range merchant {
		data = append(data, model.MerchantResponse{
			ID:        v.ID,
			Name:      v.Name,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}
	return data
}
