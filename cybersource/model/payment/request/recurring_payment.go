package request

type RecurringPaymentInformation struct {
	EndDate              string `json:"endDate,omitempty"`
	Frequency            int    `json:"frequency,omitempty"`
	NumberOfPayments     int    `json:"numberOfPayments,omitempty"`
	OriginalPurchaseDate string `json:"originalPurchaseDate,omitempty"`
	SequenceNumber       int    `json:"sequenceNumber,omitempty"`
	Type                 string `json:"type,omitempty"`
	Occurrence           string `json:"occurrence,omitempty"`
	ValidationIndicator  string `json:"validationIndicator,omitempty"`
	AmountType           string `json:"amountType,omitempty"`
	MaximumAmount        string `json:"maximumAmount,omitempty"`
	ReferenceNumber      string `json:"referenceNumber,omitempty"`
}
