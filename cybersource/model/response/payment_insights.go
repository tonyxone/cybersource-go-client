package response

type PaymentInsightsInformation struct {
	ResponseInsights *PaymentInsightsInformationResponseInsights `json:"responseInsights,omitempty"`
}

type PaymentInsightsInformationResponseInsights struct {
	Category     string `json:"category,omitempty"`
	CategoryCode string `json:"categoryCode,omitempty"`
}
