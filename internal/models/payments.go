package models

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	BaseEntity
	Date   time.Time
	Amount float32
	Method BillPaymentMethod
	UserID uuid.UUID
	BillID uuid.UUID
}

type CreatePaymentRequest struct {
	Date   time.Time         `json:"date"`
	Amount float32           `json:"amount"`
	Method BillPaymentMethod `json:"method"`
	BillID uuid.UUID         `json:"bill_id"`
	UserID uuid.UUID
}

type CreatePaymentResponse struct {
	PaymentID uuid.UUID `json:"payment_id"`
}

type PaymentItemResponse struct {
	ID     uuid.UUID         `json:"payment_id"`
	Date   time.Time         `json:"date"`
	Amount float32           `json:"amount"`
	Method BillPaymentMethod `json:"method"`
	BillID uuid.UUID         `json:"bill_id"`
	UserID uuid.UUID         `json:"user_id"`
}

type GetAllPaymentsResponse struct {
	Items []PaymentItemResponse `json:"items"`
}

type GetPaymentByIdResponse PaymentItemResponse
