package response

type TokenInformation struct {
	InstrumentIdentifierNew bool                                  `json:"instrumentidentifierNew,omitempty"`
	Customer                *TokenInformationCustomer             `json:"customer,omitempty"`
	PaymentInstrument       *TokenInformationPaymentInstrument    `json:"paymentInstrument,omitempty"`
	InstrumentIdentifier    *TokenInformationInstrumentIdentifier `json:"instrumentIdentifier,omitempty"`
	ShippingAddress         *TokenInformationShippingAddress      `json:"shippingAddress,omitempty"`
}

type TokenInformationCustomer struct {
	ID string `json:"id,omitempty"`
}

type TokenInformationPaymentInstrument struct {
	ID string `json:"id,omitempty"`
}

type TokenInformationInstrumentIdentifier struct {
	ID    string `json:"id,omitempty"`
	State string `json:"state,omitempty"`
}

type TokenInformationShippingAddress struct {
	ID string `json:"id,omitempty"`
}
