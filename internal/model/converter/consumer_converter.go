package converter

import (
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/entity"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/model"
)

func ConsumerToResponse(consumer *entity.Consumer, user *entity.User) *model.ConsumerResponse {
	return &model.ConsumerResponse{
		ID:           consumer.ID,
		User:         UserToResponse(user),
		NIK:          consumer.NIK,
		FullName:     consumer.FullName,
		LegalName:    consumer.LegalName,
		PhoneNumber:  consumer.PhoneNumber,
		Address:      consumer.Address,
		PlaceOfBirth: consumer.PlaceOfBirth,
		DateOfBirth:  consumer.DateOfBirth,
		Salary:       consumer.Salary,
		IDCardPhoto:  consumer.IDCardPhoto,
		SelfiePhoto:  consumer.SelfiePhoto,
		CreatedAt:    consumer.CreatedAt,
		UpdatedAt:    consumer.UpdatedAt,
	}
}
