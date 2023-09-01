package response

type PaymentInformation struct {
	Scheme               string                                  `json:"scheme,omitempty"`
	Bin                  string                                  `json:"bin,omitempty"`
	AccountType          string                                  `json:"accountType,omitempty"`
	Issuer               string                                  `json:"issuer,omitempty"`
	BinCountry           string                                  `json:"binCountry,omitempty"`
	Card                 *PaymentAccountInformationCard          `json:"card,omitempty"`
	TokenizedCard        *PaymentInformationTokenizedCard        `json:"tokenizedCard,omitempty"`
	AccountFeatures      *PaymentInformationAccountFeatures      `json:"accountFeatures,omitempty"`
	Bank                 *PaymentInformationBank                 `json:"bank,omitempty"`
	Customer             *PaymentInformationCustomer             `json:"token_management,omitempty"`
	PaymentInstrument    *PaymentInformationPaymentInstrument    `json:"paymentInstrument,omitempty"`
	InstrumentIdentifier *PaymentInformationInstrumentIdentifier `json:"instrumentIdentifier,omitempty"`
	ShippingAddress      *PaymentInformationShippingAddress      `json:"shippingAddress,omitempty"`
}

type PaymentInformationTokenizedCard struct {
	Prefix          string `json:"prefix,omitempty"`
	Suffix          string `json:"suffix,omitempty"`
	Type            string `json:"type,omitempty"`
	AssuranceLevel  string `json:"assuranceLevel,omitempty"`
	ExpirationMonth string `json:"expirationMonth,omitempty"`
	ExpirationYear  string `json:"expirationYear,omitempty"`
	RequestorId     string `json:"requestorId,omitempty"`
	AssuranceMethod string `json:"assuranceMethod,omitempty"`
}

type PaymentInformationAccountFeatures struct {
	AccountType        string                                       `json:"accountType,omitempty"`
	AccountStatus      string                                       `json:"accountStatus,omitempty"`
	BalanceAmount      string                                       `json:"balanceAmount,omitempty"`
	BalanceAmountType  string                                       `json:"balanceAmountType,omitempty"`
	Currency           string                                       `json:"currency,omitempty"`
	BalanceSign        string                                       `json:"balanceSign,omitempty"`
	AffluenceIndicator string                                       `json:"affluenceIndicator,omitempty"`
	Category           string                                       `json:"category,omitempty"`
	Commercial         string                                       `json:"commercial,omitempty"`
	Group              string                                       `json:"group,omitempty"`
	HealthCare         string                                       `json:"healthCare,omitempty"`
	Payroll            string                                       `json:"payroll,omitempty"`
	Level3Eligible     string                                       `json:"level3Eligible,omitempty"`
	PinlessDebit       string                                       `json:"pinlessDebit,omitempty"`
	SignatureDebit     string                                       `json:"signatureDebit,omitempty"`
	Prepaid            string                                       `json:"prepaid,omitempty"`
	Regulated          string                                       `json:"regulated,omitempty"`
	Balances           []*PaymentInformationAccountFeaturesBalances `json:"balances,omitempty"`
}

type PaymentInformationAccountFeaturesBalances struct {
	AccountType string `json:"accountType,omitempty"`
	Amount      string `json:"amount,omitempty"`
	AmountType  string `json:"amountType,omitempty"`
	Currency    string `json:"currency,omitempty"`
}

type PaymentInformationBank struct {
	Account                *PaymentInformationBankAccount `json:"account,omitempty"`
	CorrectedRoutingNumber string                         `json:"correctedRoutingNumber,omitempty"`
}

type PaymentInformationBankAccount struct {
	CorrectedAccountNumber string `json:"correctedAccountNumber,omitempty"`
}

type PaymentInformationCustomer struct {
	CustomerID string `json:"customerId,omitempty"`
	ID         string `json:"id,omitempty"`
}

type PaymentInformationPaymentInstrument struct {
	ID string `json:"id,omitempty"`
}

type PaymentInformationInstrumentIdentifier struct {
	ID    string `json:"id,omitempty"`
	State string `json:"state,omitempty"`
}

type PaymentInformationShippingAddress struct {
	ID string `json:"id,omitempty"`
}
