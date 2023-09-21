package response

type BINLookupResponse struct {
	ID                 string                              `json:"id,omitempty"`
	SubmitTimeUtc      string                              `json:"submitTimeUtc,omitempty"`
	Status             string                              `json:"status,omitempty"`
	PaymentAccountInfo *BINLookupPaymentAccountInformation `json:"paymentAccountInformation,omitempty"`
	IssuerInfo         *BINLookupIssuerInformation         `json:"issuerInformation,omitempty"`
}

type BINLookupPaymentAccountInformation struct {
	Card     *BINLookupCard     `json:"card,omitempty"`
	Features *BINLookupFeatures `json:"features,omitempty"`
}

type BINLookupCard struct {
	Type           string `json:"type,omitempty"`
	BrandName      string `json:"brandName,omitempty"`
	MaxLength      string `json:"maxLength,omitempty"`
	CredentialType string `json:"credentialType,omitempty"`
}

type BINLookupFeatures struct {
	AccountFundingSource string `json:"accountFundingSource,omitempty"`
	CardPlatform         string `json:"cardPlatform,omitempty"`
	CardProduct          string `json:"cardProduct,omitempty"`
}

type BINLookupIssuerInformation struct {
	Name          string `json:"name,omitempty"`
	Country       string `json:"country,omitempty"`
	BINLength     string `json:"BINLength,omitempty"`
	AccountPrefix string `json:"accountPrefix,omitempty"`
}
