package converter

import (
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/entity"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/model"
)

func ConsumerToResponse(consumer *entity.Consumer) *model.ConsumerResponse {
	return &model.ConsumerResponse{
		ID:           consumer.ID,
		NIK:          consumer.NIK,
		FullName:     consumer.FullName,
		LegalName:    consumer.LegalName,
		PlaceOfBirth: consumer.PlaceOfBirth,
		DateOfBirth:  consumer.DateOfBirth,
		Salary:       consumer.Salary,
		IDCardPhoto:  consumer.IDCardPhoto,
		SelfiePhoto:  consumer.SelfiePhoto,
		CreatedAt:    consumer.CreatedAt,
		UpdatedAt:    consumer.UpdatedAt,
	}
}
