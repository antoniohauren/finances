package services

import (
	"fmt"

	"github.com/antoniohauren/finances/internal/models"
	"github.com/google/uuid"
)

func (s *Services) CreatePayment(newPayment models.CreatePaymentRequest) (uuid.UUID, error) {
	dto := models.Payment{
		Date:   newPayment.Date,
		Amount: newPayment.Amount,
		Method: newPayment.Method,
		UserID: newPayment.UserID,
		BillID: newPayment.BillID,
	}

	id, err := s.repos.Payment.CreatePayment(dto)

	if err != nil {
		return uuid.Nil, err
	}

	uid, err := uuid.Parse(id)

	if err != nil {
		return uuid.Nil, err
	}

	return uid, nil
}

func (s *Services) GetPaymentByID(userID uuid.UUID, id uuid.UUID) (*models.GetPaymentByIdResponse, error) {
	payment, err := s.repos.Payment.GetPaymentByID(id)

	if err != nil {
		return nil, err
	}

	if payment.UserID != userID {
		return nil, fmt.Errorf("can't access this")
	}

	res := models.GetPaymentByIdResponse{
		ID:     payment.ID,
		Date:   payment.Date,
		Amount: payment.Amount,
		Method: payment.Method,
		UserID: payment.UserID,
	}

	return &res, nil
}

func (s *Services) GetAllPayments(userId uuid.UUID) []models.PaymentItemResponse {
	payments, err := s.repos.Payment.GetAllPayments(userId)

	if err != nil {
		return nil
	}

	items := make([]models.PaymentItemResponse, len(payments))

	for i, p := range payments {
		items[i] = models.PaymentItemResponse{
			ID:     p.ID,
			Amount: p.Amount,
			Date:   p.Date,
			Method: p.Method,
			UserID: p.UserID,
		}
	}

	return items
}

func (s *Services) GetAllPaymentsByBill(userId uuid.UUID, billId uuid.UUID) []models.PaymentItemResponse {
	payments, err := s.repos.Payment.GetAllPayments(userId)

	if err != nil {
		return nil
	}

	items := make([]models.PaymentItemResponse, len(payments))

	for i, p := range payments {
		items[i] = models.PaymentItemResponse{
			ID:     p.ID,
			Amount: p.Amount,
			Date:   p.Date,
			Method: p.Method,
			UserID: p.UserID,
			BillID: p.BillID,
		}
	}

	return items
}
