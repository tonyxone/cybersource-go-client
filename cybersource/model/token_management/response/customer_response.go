package response

import "github.com/tonyxone/cybersource-rest-sdk-go/cybersource/model/token_management"

type CustomersResponse struct {
	ID                         string                                         `json:"id,omitempty"`
	ObjectInformation          *token_management.CustomersObjectInformation   `json:"objectInformation,omitempty"`
	BuyerInformation           *token_management.BuyerInformation             `json:"buyerInformation,omitempty"`
	ClientReferenceInformation *token_management.ClientReferenceInformation   `json:"clientReferenceInformation,omitempty"`
	MerchantDefinedInformation []*token_management.MerchantDefinedInformation `json:"merchantDefinedInformation,omitempty"`
	DefaultPaymentInstrument   *token_management.DefaultPaymentInstrument     `json:"defaultPaymentInstrument,omitempty"`
	DefaultShippingAddress     *token_management.DefaultShippingAddress       `json:"defaultShippingAddress,omitempty"`
	Metadata                   *token_management.CustomersMetadata            `json:"metadata,omitempty"`
	// Reason and Message will be populated for 400 and 500 responses
	Reason  string    `json:"reason,omitempty"`
	Message string    `json:"message,omitempty"`
	Errors  []*Errors `json:"errors,omitempty"`
}

type Errors struct {
	Type    string          `json:"type,omitempty"`
	Message string          `json:"message,omitempty"`
	Details []*ErrorDetails `json:"details,omitempty"`
}

type ErrorDetails struct {
	Name     string `json:"name,omitempty"`
	Location string `json:"location,omitempty"`
}
