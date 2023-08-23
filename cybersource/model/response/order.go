package response

type OrderInformation struct {
	AmountDetails       OrderInformationAmountDetails       `json:"amountDetails,omitempty"`
	InvoiceDetails      OrderInformationInvoiceDetails      `json:"invoiceDetails,omitempty"`
	RewardPointsDetails OrderInformationRewardPointsDetails `json:"rewardPointsDetails,omitempty"`
	BillTo              OrderInformationBillTo              `json:"billTo,omitempty"`
}

type OrderInformationAmountDetails struct {
	TotalAmount      string `json:"totalAmount,omitempty"`
	AuthorizedAmount string `json:"authorizedAmount,omitempty"`
	Currency         string `json:"currency,omitempty"`
}

type OrderInformationInvoiceDetails struct {
	Level3TransmissionStatus bool `json:"level3TransmissionStatus,omitempty"`
	SalesSlipNumber          int  `json:"salesSlipNumber,omitempty"`
}

type OrderInformationRewardPointsDetails struct {
	PointsBeforeRedemption      string `json:"pointsBeforeRedemption,omitempty"`
	PointsValueBeforeRedemption string `json:"pointsValueBeforeRedemption,omitempty"`
	PointsRedeemed              string `json:"pointsRedeemed,omitempty"`
	PointsValueRedeemed         string `json:"pointsValueRedeemed,omitempty"`
	PointsAfterRedemption       string `json:"pointsAfterRedemption,omitempty"`
	PointsValueAfterRedemption  string `json:"pointsValueAfterRedemption,omitempty"`
}

type OrderInformationBillTo struct {
	AlternatePhoneNumberVerificationStatus string `json:"alternatePhoneNumberVerificationStatus,omitempty"`
	AlternateEmailVerificationStatus       string `json:"alternateEmailVerificationStatus,omitempty"`
}
