package token_management

type PaymentInstrumentBillTo struct {
	FirstName          string `json:"firstName,omitempty"`
	LastName           string `json:"lastName,omitempty"`
	Company            string `json:"company,omitempty"`
	Address1           string `json:"address1,omitempty"`
	Address2           string `json:"address2,omitempty"`
	Locality           string `json:"locality,omitempty"`
	AdministrativeArea string `json:"administrativeArea,omitempty"`
	PostalCode         string `json:"postalCode,omitempty"`
	Country            string `json:"country,omitempty"`
	Email              string `json:"email,omitempty"`
	PhoneNumber        string `json:"phoneNumber,omitempty"`
}
