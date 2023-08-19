package model

type HealthCareInformation struct {
	AmountDetails []*HealthCareInformationAmountDetails `json:"amountDetails,omitempty"`
}

type HealthCareInformationAmountDetails struct {
	AmountType string `json:"amountType,omitempty"`
	Amount     string `json:"amount,omitempty"`
}
