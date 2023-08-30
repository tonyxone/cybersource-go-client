package request

type ReverseAuthRequest struct {
	ClientReferenceInformation *ClientReferenceInformation `json:"clientReferenceInformation,omitempty"`
	ReversalInformation        *ReversalInformation        `json:"reversalInformation,omitempty"`
	ProcessingInformation      *ProcessingInformation      `json:"processingInformation,omitempty"`
	OrderInformation           *OrderInformation           `json:"orderInformation,omitempty"`
	PointOfSaleInformation     *PointOfSaleInformation     `json:"pointOfSaleInformation,omitempty"`
}

type ReversalInformation struct {
	AmountDetails *ReversalInformationAmountDetails `json:"amountDetails,omitempty"`
	Reason        string                            `json:"reason,omitempty"`
}

type ReversalInformationAmountDetails struct {
	TotalAmount string `json:"totalAmount,omitempty"`
	Currency    string `json:"currency,omitempty"`
}
