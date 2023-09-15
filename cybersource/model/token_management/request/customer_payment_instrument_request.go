package request

import "github.com/tonyxone/cybersource-rest-sdk-go/cybersource/model/token_management"

type CustomerPaymentInstrumentRequest struct {
	Default               bool                                                     `json:"default,omitempty"`
	BankAccount           *token_management.PaymentInstrumentBankAccount           `json:"bankAccount,omitempty"`
	Card                  *token_management.PaymentInstrumentCard                  `json:"card,omitempty"`
	BuyerInformation      *token_management.PaymentInstrumentBuyerInformation      `json:"buyerInformation,omitempty"`
	BillTo                *token_management.PaymentInstrumentBillTo                `json:"billTo,omitempty"`
	ProcessingInformation *token_management.PaymentInstrumentProcessingInformation `json:"processingInformation,omitempty"`
	MerchantInformation   *token_management.PaymentInstrumentMerchantInformation   `json:"merchantInformation,omitempty"`
	InstrumentIdentifier  *token_management.PaymentInstrumentIdentifier            `json:"instrumentIdentifier,omitempty"`
	Metadata              *token_management.PaymentInstrumentMetadata              `json:"metadata,omitempty"`
}
