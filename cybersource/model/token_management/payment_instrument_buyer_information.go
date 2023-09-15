package token_management

type PaymentInstrumentBuyerInformation struct {
	CompanyTaxID           string                                                     `json:"companyTaxID,omitempty"`
	Currency               string                                                     `json:"currency,omitempty"`
	DateOfBirth            string                                                     `json:"dateOfBirth,omitempty"`
	PersonalIdentification []*PaymentInstrumentBuyerInformationPersonalIdentification `json:"personalIdentification,omitempty"`
}

type PaymentInstrumentBuyerInformationPersonalIdentification struct {
	ID       string                                     `json:"id,omitempty"`
	Type     string                                     `json:"type,omitempty"`
	IssuedBy *PaymentInstrumentBuyerInformationIssuedBy `json:"issuedBy,omitempty"`
}

type PaymentInstrumentBuyerInformationIssuedBy struct {
	AdministrativeArea string `json:"administrativeArea,omitempty"`
}
