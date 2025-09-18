package models

type BillPaymentMethod string

const (
	BillPaymentMethodCreditCard BillPaymentMethod = "CREDIT_CARD"
	BillPaymentMethodDebitCard  BillPaymentMethod = "DEBIT_CARD"
	BillPaymentMethodCash       BillPaymentMethod = "CASH"
	BillPaymentMethodPix        BillPaymentMethod = "PIX"
	BillPaymentMethodOther      BillPaymentMethod = "OTHER"
)
