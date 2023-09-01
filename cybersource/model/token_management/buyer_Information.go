package token_management

type BuyerInformation struct {
	MerchantCustomerID string `json:"merchantCustomerID,omitempty"`
	Email              string `json:"email,omitempty"`
}
