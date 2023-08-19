package model

type BuyerInformation struct {
	MerchantCustomerId     string                                    `json:"merchantCustomerId,omitempty"`
	DateOfBirth            string                                    `json:"dateOfBirth,omitempty"`
	VatRegistrationNumber  string                                    `json:"vatRegistrationNumber,omitempty"`
	CompanyTaxId           string                                    `json:"companyTaxId,omitempty"`
	PersonalIdentification []*BuyerInformationPersonalIdentification `json:"personalIdentification,omitempty"`
	HashedPassword         string                                    `json:"hashedPassword,omitempty"`
	Gender                 string                                    `json:"gender,omitempty"`
	Language               string                                    `json:"language,omitempty"`
	MobilePhone            int                                       `json:"mobilePhone,omitempty"`
}

type BuyerInformationPersonalIdentification struct {
	Type                string `json:"type,omitempty"`
	ID                  string `json:"id,omitempty"`
	IssuedBy            string `json:"issuedBy,omitempty"`
	VerificationResults string `json:"verificationResults,omitempty"`
}
