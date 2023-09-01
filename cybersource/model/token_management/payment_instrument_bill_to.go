package token_management

type PaymentInstrumentBillTo struct {
	FirstName          string `json:"firstName"`
	LastName           string `json:"lastName"`
	Company            string `json:"company"`
	Address1           string `json:"address1"`
	Address2           string `json:"address2"`
	Locality           string `json:"locality"`
	AdministrativeArea string `json:"administrativeArea"`
	PostalCode         string `json:"postalCode"`
	Country            string `json:"country"`
	Email              string `json:"email"`
	PhoneNumber        string `json:"phoneNumber"`
}
