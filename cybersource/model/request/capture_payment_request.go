package request

type CapturePaymentRequest struct {
	ClientReferenceInformation *ClientReferenceInformation   `json:"clientReferenceInformation,omitempty"`
	ProcessingInformation      *ProcessingInformation        `json:"processingInformation,omitempty"`
	PaymentInformation         *PaymentInformation           `json:"paymentInformation,omitempty"`
	OrderInformation           *OrderInformation             `json:"orderInformation,omitempty"`
	BuyerInformation           *BuyerInformation             `json:"buyerInformation,omitempty"`
	DeviceInformation          *DeviceInformation            `json:"deviceInformation,omitempty"`
	MerchantInformation        *MerchantInformation          `json:"merchantInformation,omitempty"`
	AggregatorInformation      *AggregatorInformation        `json:"aggregatorInformation,omitempty"`
	PointOfSaleInformation     *PointOfSaleInformation       `json:"pointOfSaleInformation,omitempty"`
	MerchantDefinedInformation []*MerchantDefinedInformation `json:"merchantDefinedInformation,omitempty"`
	InstallmentInformation     *InstallmentInformation       `json:"installmentInformation,omitempty"`
	TravelInformation          *TravelInformation            `json:"travelInformation,omitempty"`
	PromotionInformation       *PromotionInformation         `json:"promotionInformation,omitempty"`
}
