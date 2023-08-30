package request

type MerchantInformation struct {
	MerchantDescriptor          *MerchantInformationMerchantDescriptor   `json:"merchantDescriptor,omitempty"`
	DomainName                  string                                   `json:"domainName,omitempty"`
	SalesOrganizationId         string                                   `json:"salesOrganizationId,omitempty"`
	CategoryCode                int                                      `json:"categoryCode,omitempty"`
	CategoryCodeDomestic        int                                      `json:"categoryCodeDomestic,omitempty"`
	TaxId                       string                                   `json:"taxId,omitempty"`
	VatRegistrationNumber       string                                   `json:"vatRegistrationNumber,omitempty"`
	CardAcceptorReferenceNumber string                                   `json:"cardAcceptorReferenceNumber,omitempty"`
	TransactionLocalDateTime    string                                   `json:"transactionLocalDateTime,omitempty"`
	ServiceFeeDescriptor        *MerchantInformationServiceFeeDescriptor `json:"serviceFeeDescriptor,omitempty"`
	CancelUrl                   string                                   `json:"cancelUrl,omitempty"`
	SuccessUrl                  string                                   `json:"successUrl,omitempty"`
	FailureUrl                  string                                   `json:"failureUrl,omitempty"`
	ReturnUrl                   string                                   `json:"returnUrl,omitempty"`
	PartnerIdCode               string                                   `json:"partnerIdCode,omitempty"`
	ServiceLocation             *MerchantInformationServiceLocation      `json:"serviceLocation,omitempty"`
	MerchantName                string                                   `json:"merchantName,omitempty"`
}

type MerchantInformationMerchantDescriptor struct {
	Name                       string `json:"name,omitempty"`
	AlternateName              string `json:"alternateName,omitempty"`
	Contact                    string `json:"contact,omitempty"`
	Address1                   string `json:"address1,omitempty"`
	Locality                   string `json:"locality,omitempty"`
	Country                    string `json:"country,omitempty"`
	PostalCode                 string `json:"postalCode,omitempty"`
	AdministrativeArea         string `json:"administrativeArea,omitempty"`
	Phone                      string `json:"phone,omitempty"`
	URL                        string `json:"url,omitempty"`
	CountryOfOrigin            string `json:"countryOfOrigin,omitempty"`
	CustomerServicePhoneNumber string `json:"customerServicePhoneNumber,omitempty"`
}

type MerchantInformationServiceFeeDescriptor struct {
	Name    string `json:"name,omitempty"`
	Contact string `json:"contact,omitempty"`
	State   string `json:"state,omitempty"`
}

type MerchantInformationServiceLocation struct {
	Locality               string `json:"locality,omitempty"`
	CountrySubdivisionCode string `json:"countrySubdivisionCode,omitempty"`
	CountryCode            string `json:"countryCode,omitempty"`
	PostalCode             string `json:"postalCode,omitempty"`
}
