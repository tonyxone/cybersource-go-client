package model

type CreatePaymentRequest struct {
	ClientReferenceInformation        *ClientReferenceInformation        `json:"clientReferenceInformation,omitempty"`
	ProcessingInformation             *ProcessingInformation             `json:"processingInformation,omitempty"`
	IssuerInformation                 *IssuerInformation                 `json:"issuerInformation,omitempty"`
	PaymentInformation                *PaymentInformation                `json:"paymentInformation,omitempty"`
	OrderInformation                  *OrderInformation                  `json:"orderInformation,omitempty"`
	BuyerInformation                  *BuyerInformation                  `json:"buyerInformation,omitempty"`
	RecipientInformation              *RecipientInformation              `json:"recipientInformation,omitempty"`
	DeviceInformation                 *DeviceInformation                 `json:"deviceInformation,omitempty"`
	MerchantInformation               *MerchantInformation               `json:"merchantInformation,omitempty"`
	AggregatorInformation             *AggregatorInformation             `json:"aggregatorInformation,omitempty"`
	ConsumerAuthenticationInformation *ConsumerAuthenticationInformation `json:"consumerAuthenticationInformation,omitempty"`
	PointOfSaleInformation            *PointOfSaleInformation            `json:"pointOfSaleInformation,omitempty"`
	MerchantDefinedInformation        []MerchantDefinedInformation       `json:"merchantDefinedInformation,omitempty"`
	InstallmentInformation            *InstallmentInformation            `json:"installmentInformation,omitempty"`
	TravelInformation                 *TravelInformation                 `json:"travelInformation,omitempty"`
	HealthCareInformation             *HealthCareInformation             `json:"healthCareInformation,omitempty"`
	PromotionInformation              *PromotionInformation              `json:"promotionInformation,omitempty"`
	TokenInformation                  *TokenInformation                  `json:"tokenInformation,omitempty"`
	InvoiceDetails                    *InvoiceDetails                    `json:"invoiceDetails,omitempty"`
	ProcessorInformation              *ProcessorInformation              `json:"processorInformation,omitempty"`
	RiskInformation                   *RiskInformation                   `json:"riskInformation,omitempty"`
	AcquirerInformation               *AcquirerInformation               `json:"acquirerInformation,omitempty"`
	RecurringPaymentInformation       *RecurringPaymentInformation       `json:"recurringPaymentInformation,omitempty"`
	WatchlistScreeningInformation     *WatchlistScreeningInformation     `json:"watchlistScreeningInformation,omitempty"`
}
