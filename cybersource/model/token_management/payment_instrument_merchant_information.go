package token_management

type PaymentInstrumentMerchantInformation struct {
	MerchantDescriptor *PaymentInstrumentMerchantInformationMerchantDescriptor `json:"merchantDescriptor"`
}

type PaymentInstrumentMerchantInformationMerchantDescriptor struct {
	AlternateName string `json:"alternateName"`
}
