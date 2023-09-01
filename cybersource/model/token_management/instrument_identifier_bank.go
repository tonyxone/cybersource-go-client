package token_management

type InstrumentIdentifierBankAccount struct {
	Number        string `json:"number,omitempty"`
	RoutingNumber string `json:"routingNumber,omitempty"`
}
