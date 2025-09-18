package models

type BillPaymentMethod string

const (
	BillPaymentMethodCreditCard BillPaymentMethod = "CREDIT_CARD"
	BillPaymentMethodDebitCard  BillPaymentMethod = "DEBIT_CARD"
	BillPaymentMethodCash       BillPaymentMethod = "CASH"
	BillPaymentMethodPix        BillPaymentMethod = "PIX"
	BillPaymentMethodOther      BillPaymentMethod = "OTHER"
)

func (e BillPaymentMethod) IsValid() bool {
	switch e {
	case
		BillPaymentMethodCreditCard,
		BillPaymentMethodDebitCard,
		BillPaymentMethodCash,
		BillPaymentMethodPix,
		BillPaymentMethodOther:
		return true
	default:
		return false
	}
}

type BillType string

const (
	BillTypeFixedCosts       BillType = "FIXED_COSTS"       // Custos fixos (30%)
	BillTypeComfort          BillType = "COMFORT"           // Conforto (15%)
	BillTypeGoals            BillType = "GOALS"             // Metas (15%)
	BillTypePleasures        BillType = "PLEASURES"         // Prazeres (10%)
	BillTypeFinancialFreedom BillType = "FINANCIAL_FREEDOM" // Liberdade financeira (25%)
	BillTypeKnowledge        BillType = "KNOWLEDGE"         // Conhecimento (5%)
)

func (e BillType) IsValid() bool {
	switch e {
	case
		BillTypeFixedCosts,
		BillTypeComfort,
		BillTypeGoals,
		BillTypePleasures,
		BillTypeFinancialFreedom,
		BillTypeKnowledge:
		return true
	default:
		return false
	}
}

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

func (e BillCategory) IsValid() bool {
	switch e {
	case
		BillCategoryFood,
		BillCategoryUtilities,
		BillCategoryHousing,
		BillCategoryTransport,
		BillCategoryEntertainment,
		BillCategoryHeathcare,
		BillCategorySavings,
		BillCategoryTaxes,
		BillCategoryOthers:
		return true
	default:
		return false
	}
}

type BillFrequency string

const (
	BillFrequencyYearly   BillFrequency = "YEARLY"
	BillFrequencyMontly   BillFrequency = "MONTLY"
	BillFrequencyWeekly   BillFrequency = "WEEKLY"
	BillFrequencyBiweekly BillFrequency = "BIWEEKLY"
	BillFrequencyDaily    BillFrequency = "DAILY"
	BillFrequencyOther    BillFrequency = "OTHER"
)

func (e BillFrequency) IsValid() bool {
	switch e {
	case
		BillFrequencyYearly,
		BillFrequencyMontly,
		BillFrequencyWeekly,
		BillFrequencyBiweekly,
		BillFrequencyDaily,
		BillFrequencyOther:
		return true
	default:
		return false
	}
}
