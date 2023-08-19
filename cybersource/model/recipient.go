package model

type RecipientInformation struct {
	AccountId   string `json:"accountId,omitempty"`
	LastName    string `json:"lastName,omitempty"`
	MiddleName  string `json:"middleName,omitempty"`
	PostalCode  string `json:"postalCode,omitempty"`
	DateOfBirth string `json:"dateOfBirth,omitempty"`
}
