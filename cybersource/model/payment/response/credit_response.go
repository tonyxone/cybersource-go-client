package response

type CreditResponse struct {
	ID                         string                      `json:"id,omitempty"`
	SubmitTimeUTC              string                      `json:"submitTimeUtc,omitempty"`
	Status                     string                      `json:"status,omitempty"`
	ReconciliationID           string                      `json:"reconciliationId,omitempty"`
	ClientReferenceInformation *ClientReferenceInformation `json:"clientReferenceInformation,omitempty"`
	CreditAmountDetails        *CreditAmountDetails        `json:"creditAmountDetails,omitempty"`
	ProcessingInformation      *ProcessingInformation      `json:"processingInformation,omitempty"`
	ProcessorInformation       *ProcessorInformation       `json:"processorInformation,omitempty"`
	PaymentInformation         *PaymentInformation         `json:"paymentInformation,omitempty"`
	OrderInformation           *OrderInformation           `json:"orderInformation,omitempty"`
	PointOfSaleInformation     *PointOfSaleInformation     `json:"pointOfSaleInformation,omitempty"`
	// Reason and Message will populate for 400 and 500 responses
	Reason  string `json:"reason,omitempty"`
	Message string `json:"message,omitempty"`
}

type CreditAmountDetails struct {
	CreditAmount string `json:"creditAmount,omitempty"`
	Currency     string `json:"currency,omitempty"`
}
