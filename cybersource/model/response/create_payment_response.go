package response

type CreatePaymentResponse struct {
	ID               string                                      `json:"id,omitempty"`
	Status           string                                      `json:"status,omitempty"`
	ReconciliationID string                                      `json:"reconciliationId,omitempty"`
	SubmitTimeUTC    string                                      `json:"submitTimeUtc,omitempty"`
	ErrorInformation *PaymentResponseErrorInformation            `json:"errorInformation,omitempty"`
	ClientReference  *PaymentsResponseClientReferenceInformation `json:"clientReferenceInformation,omitempty"`
}

type PaymentResponseErrorInformation struct {
	Reason  string `json:"reason,omitempty"`
	Message string `json:"message,omitempty"`
}

type PaymentsResponseErrorDetails struct {
	Field  string `json:"field,omitempty"`
	Reason string `json:"reason,omitempty"`
}

type PaymentsResponseClientReferenceInformation struct {
	Code string `json:"code"`
}
