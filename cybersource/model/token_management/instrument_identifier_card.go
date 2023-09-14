package token_management

type InstrumentIdentifierCard struct {
	Number          string `json:"number,omitempty"`
	ExpirationMonth string `json:"expirationMonth,omitempty"`
	ExpirationYear  string `json:"expirationYear,omitempty"`
	SecurityCode    string `json:"securityCode,omitempty"`
}
