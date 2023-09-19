package request

type PaymentInformation struct {
	Card                 *PaymentInformationCard                 `json:"card,omitempty"`
	TokenizedCard        *PaymentInformationTokenizedCard        `json:"tokenizedCard,omitempty"`
	FluidData            *PaymentInformationFluidData            `json:"fluidData,omitempty"`
	Customer             *PaymentInformationCustomer             `json:"token_management,omitempty"`
	PaymentInstrument    *PaymentInformationPaymentInstrument    `json:"paymentInstrument,omitempty"`
	InstrumentIdentifier *PaymentInformationInstrumentIdentifier `json:"instrumentIdentifier,omitempty"`
	ShippingAddress      *PaymentInformationShippingAddress      `json:"shippingAddress,omitempty"`
	LegacyToken          *PaymentInformationLegacyToken          `json:"legacyToken,omitempty"`
	Bank                 *PaymentInformationBank                 `json:"bank,omitempty"`
	PaymentType          *PaymentInformationPaymentType          `json:"paymentType,omitempty"`
	EWallet              *PaymentInformationEWallet              `json:"eWallet,omitempty"`
	InitiationChannel    string                                  `json:"initiationChannel,omitempty"`
}

type PaymentInformationCard struct {
	Number                   string `json:"number,omitempty"`
	ExpirationMonth          string `json:"expirationMonth,omitempty"`
	ExpirationYear           string `json:"expirationYear,omitempty"`
	Type                     string `json:"type,omitempty"`
	UseAs                    string `json:"useAs,omitempty"`
	SourceAccountType        string `json:"sourceAccountType,omitempty"`
	SourceAccountTypeDetails string `json:"sourceAccountTypeDetails,omitempty"`
	SecurityCode             string `json:"securityCode,omitempty"`
	SecurityCodeIndicator    string `json:"securityCodeIndicator,omitempty"`
	AccountEncoderId         string `json:"accountEncoderId,omitempty"`
	IssueNumber              string `json:"issueNumber,omitempty"`
	StartMonth               string `json:"startMonth,omitempty"`
	StartYear                string `json:"startYear,omitempty"`
	ProductName              string `json:"productName,omitempty"`
	TypeSelectionIndicator   string `json:"typeSelectionIndicator,omitempty"`
}

type PaymentInformationTokenizedCard struct {
	Number                string `json:"number,omitempty"`
	ExpirationMonth       string `json:"expirationMonth,omitempty"`
	ExpirationYear        string `json:"expirationYear,omitempty"`
	Type                  string `json:"type,omitempty"`
	Cryptogram            string `json:"cryptogram,omitempty"`
	RequestorId           string `json:"requestorId,omitempty"`
	TransactionType       string `json:"transactionType,omitempty"`
	AssuranceLevel        string `json:"assuranceLevel,omitempty"`
	StorageMethod         string `json:"storageMethod,omitempty"`
	SecurityCode          string `json:"securityCode,omitempty"`
	SecurityCodeIndicator string `json:"securityCodeIndicator,omitempty"`
	AssuranceMethod       string `json:"assuranceMethod,omitempty"`
}

type PaymentInformationFluidData struct {
	KeySerialNumber string `json:"keySerialNumber,omitempty"`
	Descriptor      string `json:"descriptor,omitempty"`
	Value           string `json:"value,omitempty"`
	Encoding        string `json:"encoding,omitempty"`
}

type PaymentInformationCustomer struct {
	CustomerID string `json:"customerId,omitempty"`
	ID         string `json:"id,omitempty"`
}

type PaymentInformationPaymentInstrument struct {
	ID string `json:"id,omitempty"`
}

type PaymentInformationInstrumentIdentifier struct {
	ID string `json:"id,omitempty"`
}

type PaymentInformationShippingAddress struct {
	ID string `json:"id,omitempty"`
}

type PaymentInformationLegacyToken struct {
	ID string `json:"id,omitempty"`
}

type PaymentInformationBank struct {
	Account       *PaymentInformationBankAccount `json:"account,omitempty"`
	RoutingNumber string                         `json:"routingNumber,omitempty"`
	Iban          string                         `json:"iban,omitempty"`
	SwiftCode     string                         `json:"swiftCode,omitempty"`
	Code          string                         `json:"code,omitempty"`
}

type PaymentInformationBankAccount struct {
	Type                      string `json:"type,omitempty"`
	Number                    string `json:"number,omitempty"`
	EncoderId                 string `json:"encoderId,omitempty"`
	CheckNumber               string `json:"checkNumber,omitempty"`
	CheckImageReferenceNumber string `json:"checkImageReferenceNumber,omitempty"`
}

type PaymentInformationPaymentType struct {
	Name        string `json:"name,omitempty"`
	SubTypeName string `json:"subTypeName,omitempty"`
	Type        string `json:"type,omitempty"`
}

type PaymentInformationPaymentTypeMethod struct {
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}

type PaymentInformationEWallet struct {
	AccountId string `json:"accountId,omitempty"`
}
