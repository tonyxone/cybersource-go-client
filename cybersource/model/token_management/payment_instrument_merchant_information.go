package token_management

type PaymentInstrumentMerchantInformation struct {
	MerchantDescriptor *PaymentInstrumentMerchantInformationMerchantDescriptor `json:"merchantDescriptor,omitempty"`
}

type PaymentInstrumentMerchantInformationMerchantDescriptor struct {
	AlternateName string `json:"alternateName,omitempty"`
}
