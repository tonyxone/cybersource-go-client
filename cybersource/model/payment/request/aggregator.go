package request

type AggregatorInformation struct {
	AggregatorID string                            `json:"aggregatorId,omitempty"`
	Name         string                            `json:"name,omitempty"`
	SubMerchant  *AggregatorInformationSubMerchant `json:"subMerchant,omitempty"`
}

type AggregatorInformationSubMerchant struct {
	CardAcceptorID     string `json:"cardAcceptorId,omitempty"`
	ID                 string `json:"id,omitempty"`
	Name               string `json:"name,omitempty"`
	Address1           string `json:"address1,omitempty"`
	Locality           string `json:"locality,omitempty"`
	AdministrativeArea string `json:"administrativeArea,omitempty"`
	Region             string `json:"region,omitempty"`
	PostalCode         string `json:"postalCode,omitempty"`
	Country            string `json:"country,omitempty"`
	Email              string `json:"email,omitempty"`
	PhoneNumber        string `json:"phoneNumber,omitempty"`
}
