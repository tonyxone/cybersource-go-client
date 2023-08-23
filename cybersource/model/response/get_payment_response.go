package response

type GetPaymentResponse struct {
	ID                         string                      `json:"id,omitempty"`
	ReconciliationID           string                      `json:"reconciliationId,omitempty"`
	SubmitTimeUTC              string                      `json:"submitTimeUtc,omitempty"`
	ProcessorInformation       *ProcessorInformation       `json:"processorInformation,omitempty"`
	PaymentInformation         *PaymentInformation         `json:"paymentInformation,omitempty"`
	OrderInformation           *OrderInformation           `json:"orderInformation,omitempty"`
	ClientReferenceInformation *ClientReferenceInformation `json:"clientReferenceInformation,omitempty"`
	StatusInformation          *StatusInformation          `json:"statusInformation,omitempty"`
}

type StatusInformation struct {
	Reason  string `json:"reason,omitempty"`
	Message string `json:"message,omitempty"`
}
