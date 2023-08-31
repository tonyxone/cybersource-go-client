package token_management

type InstrumentIdentifierTokenizedCard struct {
	Type            string                `json:"type,omitempty"`
	State           string                `json:"state,omitempty"`
	Reason          string                `json:"reason,omitempty"`
	Number          string                `json:"number,omitempty"`
	ExpirationMonth string                `json:"expirationMonth,omitempty"`
	ExpirationYear  string                `json:"expirationYear,omitempty"`
	Cryptogram      string                `json:"cryptogram,omitempty"`
	Card            *TokenizedCardDetails `json:"card,omitempty"`
}

type TokenizedCardDetails struct {
	Suffix          string `json:"suffix,omitempty"`
	ExpirationMonth string `json:"expirationMonth,omitempty"`
	ExpirationYear  string `json:"expirationYear,omitempty"`
}
