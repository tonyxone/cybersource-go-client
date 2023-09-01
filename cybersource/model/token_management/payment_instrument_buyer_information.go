package token_management

type PaymentInstrumentBuyerInformation struct {
	CompanyTaxID           string                                                     `json:"companyTaxID"`
	Currency               string                                                     `json:"currency"`
	DateOfBirth            string                                                     `json:"dateOfBirth"`
	PersonalIdentification []*PaymentInstrumentBuyerInformationPersonalIdentification `json:"personalIdentification"`
}

type PaymentInstrumentBuyerInformationPersonalIdentification struct {
	ID       string                                     `json:"id"`
	Type     string                                     `json:"type"`
	IssuedBy *PaymentInstrumentBuyerInformationIssuedBy `json:"issuedBy"`
}

type PaymentInstrumentBuyerInformationIssuedBy struct {
	AdministrativeArea string `json:"administrativeArea"`
}
