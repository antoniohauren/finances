package models

import (
	"time"

	"github.com/google/uuid"
)

type Bill struct {
	BaseEntity
	Name          string
	DueDate       time.Time
	Type          BillType
	Category      BillCategory
	Frequency     BillFrequency
	PaymentMethod BillPaymentMethod
	UserID        uuid.UUID
}

type CreateBillRequest struct {
	Name          string            `json:"name"`
	DueDate       time.Time         `json:"due_date"`
	Type          BillType          `json:"type"`
	Category      BillCategory      `json:"category"`
	Frequency     BillFrequency     `json:"frequency"`
	PaymentMethod BillPaymentMethod `json:"payment_method"`
	UserID        uuid.UUID
}

type CreateBillResponse struct {
	BillID uuid.UUID `json:"bill_id"`
}

type BillItemResponse struct {
	ID            uuid.UUID         `json:"bill_id"`
	Name          string            `json:"name"`
	DueDate       time.Time         `json:"due_date"`
	Type          BillType          `json:"type"`
	Category      BillCategory      `json:"category"`
	Frequency     BillFrequency     `json:"frequency"`
	PaymentMethod BillPaymentMethod `json:"payment_method"`
	UserID        uuid.UUID         `json:"user_id"`
}

type GetAllBillsResponse struct {
	Items []BillItemResponse `json:"items"`
}

type GetBillByIDResponse BillItemResponse
