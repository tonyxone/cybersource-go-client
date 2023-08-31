package token_management

type PaymentInstrumentCard struct {
	ExpirationMonth      string                                     `json:"expirationMonth"`
	ExpirationYear       string                                     `json:"expirationYear"`
	Type                 string                                     `json:"type"`
	IssueNumber          string                                     `json:"issueNumber"`
	StartMonth           string                                     `json:"startMonth"`
	StartYear            string                                     `json:"startYear"`
	UseAs                string                                     `json:"useAs"`
	Hash                 string                                     `json:"hash"`
	TokenizedInformation *PaymentInstrumentCardTokenizedInformation `json:"tokenizedInformation"`
}

type PaymentInstrumentCardTokenizedInformation struct {
	RequestorID     string `json:"requestorID"`
	TransactionType string `json:"transactionType"`
}
