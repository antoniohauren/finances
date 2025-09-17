package services

import (
	"fmt"

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

func (s *Services) GetAllBills(userID uuid.UUID) []models.BillItemResponse {
	bills, err := s.repos.Bill.GetAllBills(userID)

	if err != nil {
		return nil
	}

	items := make([]models.BillItemResponse, len(bills))

	for i, b := range bills {
		items[i] = models.BillItemResponse{
			ID:            b.ID,
			Name:          b.Name,
			DueDate:       b.DueDate,
			Type:          b.Type,
			Category:      b.Category,
			Frequency:     b.Frequency,
			PaymentMethod: b.PaymentMethod,
			UserID:        b.UserID,
		}
	}

	return items
}

func (s *Services) GetBillByID(userID uuid.UUID, id uuid.UUID) (*models.GetBillByIDResponse, error) {
	bill, err := s.repos.Bill.GetBillByID(id)

	if err != nil {
		return nil, err
	}

	if bill.UserID != userID {
		return nil, fmt.Errorf("can't access this bill")
	}

	res := models.GetBillByIDResponse{
		ID:            bill.ID,
		Name:          bill.Name,
		DueDate:       bill.DueDate,
		Type:          bill.Type,
		Category:      bill.Category,
		Frequency:     bill.Frequency,
		PaymentMethod: bill.PaymentMethod,
		UserID:        bill.UserID,
	}

	return &res, nil
}
