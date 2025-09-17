package services

import (
	"github.com/antoniohauren/finances/internal/models"
	"github.com/google/uuid"
)

func (s *Services) CreateBill(newBill models.CreateBillRequest) (uuid.UUID, error) {
	dto := models.Bill{
		Name:          newBill.Name,
		DueDate:       newBill.DueDate,
		Type:          newBill.Type,
		Category:      newBill.Category,
		Frequency:     newBill.Frequency,
		PaymentMethod: newBill.PaymentMethod,
		UserID:        newBill.UserID,
	}

	id, err := s.repos.Bill.CreateBill(dto)

	if err != nil {
		return uuid.Nil, err
	}

	uid, err := uuid.Parse(id)

	if err != nil {
		return uuid.Nil, err
	}

	return uid, nil
}
