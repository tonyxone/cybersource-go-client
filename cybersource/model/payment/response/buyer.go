package response

type BuyerInformation struct {
	MerchantCustomerId     string                                    `json:"merchantCustomerId,omitempty"`
	DateOfBirth            string                                    `json:"dateOfBirth,omitempty"`
	VatRegistrationNumber  string                                    `json:"vatRegistrationNumber,omitempty"`
	PersonalIdentification []*BuyerInformationPersonalIdentification `json:"personalIdentification,omitempty"`
	TaxId                  string                                    `json:"taxId,omitempty"`
}

type BuyerInformationPersonalIdentification struct {
	Type                string `json:"type,omitempty"`
	ID                  string `json:"id,omitempty"`
	IssuedBy            string `json:"issuedBy,omitempty"`
	VerificationResults string `json:"verificationResults,omitempty"`
}
