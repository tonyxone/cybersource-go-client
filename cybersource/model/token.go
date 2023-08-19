package model

type TokenInformation struct {
	Jti                string                             `json:"jti,omitempty"`
	TransientTokenJwt  string                             `json:"transientTokenJwt,omitempty"`
	PaymentInstrument  *TokenInformationPaymentInstrument `json:"paymentInstrument,omitempty"`
	ShippingAddress    *TokenInformationShippingAddress   `json:"shippingAddress,omitempty"`
	NetworkTokenOption string                             `json:"networkTokenOption,omitempty"`
}

type TokenInformationPaymentInstrument struct {
	Default bool `json:"default,omitempty"`
}

type TokenInformationShippingAddress struct {
	Default bool `json:"default,omitempty"`
}
