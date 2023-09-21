package request

type BINLookupRequest struct {
	PaymentInformation *BINLookupPaymentInformation `json:"paymentInformation,omitempty"`
	TokenInformation   *BINLookupTokenInformation   `json:"tokenInformation,omitempty"`
}

type BINLookupPaymentInformation struct {
	Card                 *BINLookupCard                 `json:"card,omitempty"`
	Customer             *BINLookupCustomer             `json:"customer,omitempty"`
	InstrumentIdentifier *BINLookupInstrumentIdentifier `json:"instrumentIdentifier,omitempty"`
	PaymentInstrument    *BINLookupPaymentInstrument    `json:"paymentInstrument,omitempty"`
}

type BINLookupCard struct {
	Number string `json:"number,omitempty"`
}

type BINLookupCustomer struct {
	ID string `json:"id,omitempty"`
}

type BINLookupInstrumentIdentifier struct {
	ID string `json:"id,omitempty"`
}

type BINLookupPaymentInstrument struct {
	ID string `json:"id,omitempty"`
}

type BINLookupTokenInformation struct {
	Jti               string `json:"jti,omitempty"`
	TransientTokenJwt string `json:"transientTokenJwt,omitempty"`
}
