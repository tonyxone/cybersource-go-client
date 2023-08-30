package request

import "github.com/tonyxone/cybersource-rest-sdk-go/cybersource/model/customer"

type CustomerRequest struct {
	ID                         string                                 `json:"id"`
	ObjectInformation          *customer.CustomersObjectInformation   `json:"objectInformation"`
	BuyerInformation           *customer.BuyerInformation             `json:"buyerInformation"`
	ClientReferenceInformation *customer.ClientReferenceInformation   `json:"clientReferenceInformation"`
	MerchantDefinedInformation []*customer.MerchantDefinedInformation `json:"merchantDefinedInformation"`
	DefaultPaymentInstrument   *customer.DefaultPaymentInstrument     `json:"defaultPaymentInstrument"`
	DefaultShippingAddress     *customer.DefaultShippingAddress       `json:"defaultShippingAddress"`
	Metadata                   *customer.CustomersMetadata            `json:"metadata"`
}
