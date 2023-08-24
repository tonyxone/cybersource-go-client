package request

type VoidRequest struct {
	ClientReferenceInformation *ClientReferenceInformation `json:"clientReferenceInformation,omitempty"`
	PaymentInformation         *PaymentInformation         `json:"paymentInformation,omitempty"`
	OrderInformation           *OrderInformation           `json:"orderInformation,omitempty"`
}
