package request

import "github.com/tonyxone/cybersource-rest-sdk-go/cybersource/model"

type CreatePaymentRequest struct {
	ClientReferenceInformation        *model.ClientReferenceInformation        `json:"clientReferenceInformation,omitempty"`
	ProcessingInformation             *model.ProcessingInformation             `json:"processingInformation,omitempty"`
	IssuerInformation                 *model.IssuerInformation                 `json:"issuerInformation,omitempty"`
	PaymentInformation                *model.PaymentInformation                `json:"paymentInformation,omitempty"`
	OrderInformation                  *model.OrderInformation                  `json:"orderInformation,omitempty"`
	BuyerInformation                  *model.BuyerInformation                  `json:"buyerInformation,omitempty"`
	RecipientInformation              *model.RecipientInformation              `json:"recipientInformation,omitempty"`
	DeviceInformation                 *model.DeviceInformation                 `json:"deviceInformation,omitempty"`
	MerchantInformation               *model.MerchantInformation               `json:"merchantInformation,omitempty"`
	AggregatorInformation             *model.AggregatorInformation             `json:"aggregatorInformation,omitempty"`
	ConsumerAuthenticationInformation *model.ConsumerAuthenticationInformation `json:"consumerAuthenticationInformation,omitempty"`
	PointOfSaleInformation            *model.PointOfSaleInformation            `json:"pointOfSaleInformation,omitempty"`
	MerchantDefinedInformation        []*model.MerchantDefinedInformation      `json:"merchantDefinedInformation,omitempty"`
	InstallmentInformation            *model.InstallmentInformation            `json:"installmentInformation,omitempty"`
	TravelInformation                 *model.TravelInformation                 `json:"travelInformation,omitempty"`
	HealthCareInformation             *model.HealthCareInformation             `json:"healthCareInformation,omitempty"`
	PromotionInformation              *model.PromotionInformation              `json:"promotionInformation,omitempty"`
	TokenInformation                  *model.TokenInformation                  `json:"tokenInformation,omitempty"`
	InvoiceDetails                    *model.InvoiceDetails                    `json:"invoiceDetails,omitempty"`
	ProcessorInformation              *model.ProcessorInformation              `json:"processorInformation,omitempty"`
	RiskInformation                   *model.RiskInformation                   `json:"riskInformation,omitempty"`
	AcquirerInformation               *model.AcquirerInformation               `json:"acquirerInformation,omitempty"`
	RecurringPaymentInformation       *model.RecurringPaymentInformation       `json:"recurringPaymentInformation,omitempty"`
	WatchlistScreeningInformation     *model.WatchlistScreeningInformation     `json:"watchlistScreeningInformation,omitempty"`
}
