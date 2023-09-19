package response

import "github.com/tonyxone/cybersource-rest-sdk-go/cybersource/model/payment/request"

type TransactionDetails struct {
	ID                                string                                     `json:"id,omitempty"`
	RootID                            string                                     `json:"rootId,omitempty"`
	ReconciliationID                  string                                     `json:"reconciliationId,omitempty"`
	MerchantID                        string                                     `json:"merchantId,omitempty"`
	SubmitTimeUTC                     string                                     `json:"submitTimeUTC,omitempty"`
	Status                            string                                     `json:"status,omitempty"`
	ApplicationInformation            *ApplicationInformation                    `json:"applicationInformation,omitempty"`
	BuyerInformation                  *request.BuyerInformation                  `json:"buyerInformation,omitempty"`
	ClientReferenceInformation        *request.ClientReferenceInformation        `json:"clientReferenceInformation,omitempty"`
	ConsumerAuthenticationInformation *request.ConsumerAuthenticationInformation `json:"consumerAuthenticationInformation,omitempty"`
	DeviceInformation                 *request.DeviceInformation                 `json:"deviceInformation,omitempty"`
	ErrorInformation                  *ErrorInformation                          `json:"errorInformation,omitempty"`
	InstallmentInformation            *request.InstallmentInformation            `json:"installmentInformation,omitempty"`
	FraudMarkingInformation           *FraudMarkingInformation                   `json:"fraudMarkingInformation,omitempty"`
	HealthCareInformation             *request.HealthCareInformation             `json:"healthCareInformation,omitempty"`
	MerchantDefinedInformation        []*request.MerchantDefinedInformation      `json:"merchantDefinedInformation,omitempty"`
	MerchantInformation               *request.MerchantInformation               `json:"merchantInformation,omitempty"`
	OrderInformation                  *request.OrderInformation                  `json:"orderInformation,omitempty"`
	PaymentInformation                *request.PaymentInformation                `json:"paymentInformation,omitempty"`
	PaymentInsightsInformation        *PaymentInsightsInformation                `json:"paymentInsightsInformation,omitempty"`
	PayoutOptions                     *PayoutOptions                             `json:"payoutOptions,omitempty"`
	ProcessingInformation             *request.ProcessingInformation             `json:"processingInformation,omitempty"`
	ProcessorInformation              *request.ProcessorInformation              `json:"processorInformation,omitempty"`
	PointOfSaleInformation            *request.PointOfSaleInformation            `json:"pointOfSaleInformation,omitempty"`
	RiskInformation                   *request.RiskInformation                   `json:"riskInformation,omitempty"`
	SenderInformation                 *SenderInformation                         `json:"senderInformation,omitempty"`
	TokenInformation                  *request.TokenInformation                  `json:"tokenInformation,omitempty"`
}

type ApplicationInformation struct {
	Status       string          `json:"status,omitempty"`
	ReasonCode   int             `json:"reasonCode,omitempty"`
	RCode        string          `json:"rCode,omitempty"`
	RFlag        string          `json:"rFlag,omitempty"`
	Applications []*Applications `json:"applications,omitempty"`
}

type Applications struct {
	Name             string `json:"name,omitempty"`
	Status           string `json:"status,omitempty"`
	ReasonCode       string `json:"reasonCode,omitempty"`
	RCode            string `json:"rCode,omitempty"`
	RFlag            string `json:"rFlag,omitempty"`
	ReconciliationID string `json:"reconciliationId,omitempty"`
	RMessage         string `json:"rMessage,omitempty"`
	ReturnCode       int    `json:"returnCode,omitempty"`
}

type FraudMarkingInformation struct {
	Reason string `json:"reason,omitempty"`
}

type PayoutOptions struct {
	PayoutInquiry string `json:"payoutInquiry,omitempty"`
}

type SenderInformation struct {
	ReferenceNumber string `json:"referenceNumber,omitempty"`
}
