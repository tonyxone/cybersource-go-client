package token_management

type PaymentInstrumentProcessingInformation struct {
	BillPaymentProgramEnabled bool                                                       `json:"billPaymentProgramEnabled,omitempty"`
	BankTransferOptions       *PaymentInstrumentProcessingInformationBankTransferOptions `json:"bankTransferOptions,omitempty"`
}

type PaymentInstrumentProcessingInformationBankTransferOptions struct {
	SECCode string `json:"SECCode,omitempty"`
}
