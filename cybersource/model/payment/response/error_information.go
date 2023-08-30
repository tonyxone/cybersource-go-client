package response

type ErrorInformation struct {
	Reason  string                          `json:"reason,omitempty"`
	Message string                          `json:"message,omitempty"`
	Details []*PaymentsResponseErrorDetails `json:"details,omitempty"`
}

type PaymentsResponseErrorDetails struct {
	Field  string `json:"field,omitempty"`
	Reason string `json:"reason,omitempty"`
}
