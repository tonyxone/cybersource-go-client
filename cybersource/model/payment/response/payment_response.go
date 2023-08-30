package response

type PaymentResponse struct {
	ID                                string                             `json:"id,omitempty"`
	Status                            string                             `json:"status,omitempty"`
	ReconciliationID                  string                             `json:"reconciliationId,omitempty"`
	SubmitTimeUTC                     string                             `json:"submitTimeUtc,omitempty"`
	ErrorInformation                  *ErrorInformation                  `json:"errorInformation,omitempty"`
	ClientReference                   *ClientReferenceInformation        `json:"clientReferenceInformation,omitempty"`
	ProcessingInformation             *ProcessingInformation             `json:"processingInformation,omitempty"`
	IssuerInformation                 *IssuerInformation                 `json:"issuerInformation,omitempty"`
	PaymentAccountInformation         *PaymentAccountInformation         `json:"paymentAccountInformation,omitempty"`
	PaymentInformation                *PaymentInformation                `json:"paymentInformation,omitempty"`
	PaymentInsightsInformation        *PaymentInsightsInformation        `json:"paymentInsightsInformation,omitempty"`
	OrderInformation                  *OrderInformation                  `json:"orderInformation,omitempty"`
	PointOfSaleInformation            *PointOfSaleInformation            `json:"pointOfSaleInformation,omitempty"`
	InstallmentInformation            *InstallmentInformation            `json:"installmentInformation,omitempty"`
	TokenInformation                  *TokenInformation                  `json:"tokenInformation,omitempty"`
	BuyerInformation                  *BuyerInformation                  `json:"buyerInformation,omitempty"`
	RiskInformation                   *RiskInformation                   `json:"riskInformation,omitempty"`
	ConsumerAuthenticationInformation *ConsumerAuthenticationInformation `json:"consumerAuthenticationInformation,omitempty"`
	EmbeddedActions                   *EmbeddedActions                   `json:"embeddedActions,omitempty"`
	WatchlistScreeningInformation     *WatchlistScreeningInformation     `json:"watchlistScreeningInformation,omitempty"`
	// Reason and Message will populate for 400 and 500 responses
	Reason  string `json:"reason,omitempty"`
	Message string `json:"message,omitempty"`
}
