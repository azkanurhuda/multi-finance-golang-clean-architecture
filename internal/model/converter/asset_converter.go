package converter

import (
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/entity"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/model"
)

func AssetToResponse(asset *entity.Asset, merchant *entity.Merchant) *model.AssetResponse {
	return &model.AssetResponse{
		ID:          asset.ID,
		Merchant:    MerchantToResponse(merchant),
		Name:        asset.Name,
		OTR:         asset.OTR,
		AdminFee:    asset.AdminFee,
		Description: asset.Description,
		CreatedAt:   asset.CreatedAt,
		UpdatedAt:   asset.UpdatedAt,
	}
}

func ListAssetToResponse(asset []entity.Asset) []model.AssetResponse {
	var data []model.AssetResponse
	for _, v := range asset {
		data = append(data, model.AssetResponse{
			ID:          v.ID,
			Merchant:    nil,
			Name:        v.Name,
			OTR:         v.OTR,
			AdminFee:    v.AdminFee,
			Description: v.Description,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
		})
	}
	return data
}
