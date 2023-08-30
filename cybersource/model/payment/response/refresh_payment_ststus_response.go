package response

type RefreshPaymentStatusResponse struct {
	ID                         string                      `json:"id,omitempty"`
	Status                     string                      `json:"status,omitempty"`
	ReconciliationID           string                      `json:"reconciliationId,omitempty"`
	SubmitTimeUTC              string                      `json:"submitTimeUtc,omitempty"`
	ProcessorInformation       *ProcessorInformation       `json:"processorInformation,omitempty"`
	PaymentInformation         *PaymentInformation         `json:"paymentInformation,omitempty"`
	OrderInformation           *OrderInformation           `json:"orderInformation,omitempty"`
	ClientReferenceInformation *ClientReferenceInformation `json:"clientReferenceInformation,omitempty"`
	// Reason and Message will populate for 400 and 500 responses
	Reason  string `json:"reason,omitempty"`
	Message string `json:"message,omitempty"`
}
