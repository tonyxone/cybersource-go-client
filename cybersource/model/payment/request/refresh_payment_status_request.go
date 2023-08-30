package request

type RefreshPaymentStatusRequest struct {
	PaymentInformation         *PaymentInformation         `json:"paymentInformation,omitempty"`
	ClientReferenceInformation *ClientReferenceInformation `json:"clientReferenceInformation,omitempty"`
	ProcessingInformation      *ProcessingInformation      `json:"processingInformation,omitempty"`
	AgreementInformation       *AgreementInformation       `json:"agreementInformation,omitempty"`
}
