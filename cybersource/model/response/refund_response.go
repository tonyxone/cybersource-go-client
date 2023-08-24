package response

type RefundResponse struct {
	ID                         string                      `json:"id,omitempty"`
	SubmitTimeUtc              string                      `json:"submitTimeUtc,omitempty"`
	Status                     string                      `json:"status,omitempty"`
	ReconciliationID           string                      `json:"reconciliationId,omitempty"`
	ClientReferenceInformation *ClientReferenceInformation `json:"clientReferenceInformation,omitempty"`
	RefundAmountDetails        *RefundAmountDetails        `json:"refundAmountDetails,omitempty"`
	ProcessorInformation       *ProcessorInformation       `json:"processorInformation,omitempty"`
	OrderInformation           *OrderInformation           `json:"orderInformation,omitempty"`
	PointOfSaleInformation     *PointOfSaleInformation     `json:"pointOfSaleInformation,omitempty"`
}

type RefundAmountDetails struct {
	RefundAmount string `json:"refundAmount,omitempty"`
	CreditAmount string `json:"creditAmount,omitempty"`
	Currency     string `json:"currency,omitempty"`
}
