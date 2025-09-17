package models

import (
	"time"

	"github.com/google/uuid"
)

type BillType string

const (
	BillTypeFixedCosts       BillType = "FIXED_COSTS"       // Custos fixos (30%)
	BillTypeComfort          BillType = "COMFORT"           // Conforto (15%)
	BillTypeGoals            BillType = "GOALS"             // Metas (15%)
	BillTypePleasures        BillType = "PLEASURES"         // Prazeres (10%)
	BillTypeFinancialFreedom BillType = "FINANCIAL_FREEDOM" // Liberdade financeira (25%)
	BillTypeKnowledge        BillType = "KNOWLEDGE"         // Conhecimento (5%)
)

type BillCategory string

const (
	BillCategoryFood          BillCategory = "FOOD"
	BillCategoryUtilities     BillCategory = "UTILITIES"
	BillCategoryHousing       BillCategory = "HOUSING"
	BillCategoryTransport     BillCategory = "TRANSPORT"
	BillCategoryEntertainment BillCategory = "ENTERTAINMENT"
	BillCategoryHeathcare     BillCategory = "HEALTHCARE"
	BillCategorySavings       BillCategory = "SAVINGS"
	BillCategoryTaxes         BillCategory = "TAXES"
	BillCategoryOthers        BillCategory = "OTHERS"
)

type BillFrequency string

const (
	BillFrequencyYearly   BillFrequency = "YEARLY"
	BillFrequencyMontly   BillFrequency = "MONTLY"
	BillFrequencyWeekly   BillFrequency = "WEEKLY"
	BillFrequencyBiweekly BillFrequency = "BIWEEKLY"
	BillFrequencyDaily    BillFrequency = "DAILY"
	BillFrequencyOther    BillFrequency = "OTHER"
)

type BillPaymentMethod string

const (
	BillPaymentMethodCreditCard BillPaymentMethod = "CREDIT_CARD"
	BillPaymentMethodDebitCard  BillPaymentMethod = "DEBIT_CARD"
	BillPaymentMethodCash       BillPaymentMethod = "CASH"
	BillPaymentMethodPix        BillPaymentMethod = "PIX"
	BillPaymentMethodOther      BillPaymentMethod = "OTHER"
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
