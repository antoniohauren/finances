package models

import (
	"github.com/google/uuid"
)

type Report struct {
	BillID        uuid.UUID
	BillName      string
	Frequency     string
	BillMethod    string
	TotalAmount   float64
	PaymentsCount int
}

type ReportItemResponse struct {
	BillID        uuid.UUID `json:"bill_id"`
	BillName      string    `json:"bill_name"`
	Frequency     string    `json:"frequency"`
	BillMethod    string    `json:"bill_method"`
	TotalAmount   float64   `json:"total_amount"`
	PaymentsCount int       `json:"payments_count"`
}

type ReportResponse struct {
	Items []ReportItemResponse `json:"items"`
}
