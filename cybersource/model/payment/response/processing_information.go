package response

type ProcessingInformation struct {
	BankTransferOptions *BankTransferOptions `json:"bankTransferOptions,omitempty"`
	PaymentSolution     string               `json:"paymentSolution,omitempty"`
	EnhancedDataEnabled *bool                `json:"enhancedDataEnabled,omitempty"`
}

type BankTransferOptions struct {
	SettlementMethod    string `json:"settlementMethod,omitempty"`
	FraudScreeningLevel string `json:"fraudScreeningLevel,omitempty"`
}
