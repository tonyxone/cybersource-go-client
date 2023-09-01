package request

import "github.com/tonyxone/cybersource-rest-sdk-go/cybersource/model/token_management"

type CustomerPaymentInstrumentRequest struct {
	BankAccount           *token_management.PaymentInstrumentBankAccount           `json:"bankAccount"`
	Card                  *token_management.PaymentInstrumentCard                  `json:"card"`
	BuyerInformation      *token_management.PaymentInstrumentBuyerInformation      `json:"buyerInformation"`
	BillTo                *token_management.PaymentInstrumentBillTo                `json:"billTo"`
	ProcessingInformation *token_management.PaymentInstrumentProcessingInformation `json:"processingInformation"`
	MerchantInformation   *token_management.PaymentInstrumentMerchantInformation   `json:"merchantInformation"`
	InstrumentIdentifier  *token_management.PaymentInstrumentIdentifier            `json:"instrumentIdentifier"`
	Metadata              *token_management.PaymentInstrumentMetadata              `json:"metadata"`
}
