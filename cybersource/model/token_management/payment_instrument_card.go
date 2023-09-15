package token_management

type PaymentInstrumentCard struct {
	ExpirationMonth      string                                     `json:"expirationMonth,omitempty"`
	ExpirationYear       string                                     `json:"expirationYear,omitempty"`
	Type                 string                                     `json:"type,omitempty"`
	IssueNumber          string                                     `json:"issueNumber,omitempty"`
	StartMonth           string                                     `json:"startMonth,omitempty"`
	StartYear            string                                     `json:"startYear,omitempty"`
	UseAs                string                                     `json:"useAs,omitempty"`
	Hash                 string                                     `json:"hash,omitempty"`
	TokenizedInformation *PaymentInstrumentCardTokenizedInformation `json:"tokenizedInformation,omitempty"`
}

type PaymentInstrumentCardTokenizedInformation struct {
	RequestorID     string `json:"requestorID,omitempty"`
	TransactionType string `json:"transactionType,omitempty"`
}
