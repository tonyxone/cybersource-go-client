package token_management

type PaymentInstrumentProcessingInformation struct {
	BillPaymentProgramEnabled bool                                                       `json:"billPaymentProgramEnabled"`
	BankTransferOptions       *PaymentInstrumentProcessingInformationBankTransferOptions `json:"bankTransferOptions"`
}

type PaymentInstrumentProcessingInformationBankTransferOptions struct {
	SECCode string `json:"SECCode"`
}
