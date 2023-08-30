package response

type PaymentAccountInformation struct {
	Card *PaymentAccountInformationCard `json:"card,omitempty"`
}

type PaymentAccountInformationCard struct {
	Suffix          string `json:"suffix,omitempty"`
	ExpirationMonth string `json:"expirationMonth,omitempty"`
	ExpirationYear  string `json:"expirationYear,omitempty"`
	Type            string `json:"type,omitempty"`
	Prefix          string `json:"prefix,omitempty"`
	HashedNumber    string `json:"hashedNumber,omitempty"`
}
