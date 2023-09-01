package response

type ProcessorInformation struct {
	AuthIndicator                  string                                              `json:"authIndicator,omitempty"`
	ApprovalCode                   string                                              `json:"approvalCode,omitempty"`
	CardReferenceData              string                                              `json:"cardReferenceData,omitempty"`
	TransactionId                  string                                              `json:"transactionId,omitempty"`
	NetworkTransactionId           string                                              `json:"networkTransactionId,omitempty"`
	ResponseCode                   string                                              `json:"responseCode,omitempty"`
	ResponseCodeSource             string                                              `json:"responseCodeSource,omitempty"`
	ResponseDetails                string                                              `json:"responseDetails,omitempty"`
	ResponseCategoryCode           string                                              `json:"responseCategoryCode,omitempty"`
	ForwardedAcquirerCode          string                                              `json:"forwardedAcquirerCode,omitempty"`
	SettlementDate                 string                                              `json:"settlementDate,omitempty"`
	SystemTraceAuditNumber         string                                              `json:"systemTraceAuditNumber,omitempty"`
	PaymentAccountReferenceNumber  string                                              `json:"paymentAccountReferenceNumber,omitempty"`
	TransactionIntegrityCode       string                                              `json:"transactionIntegrityCode,omitempty"`
	AmexVerbalAuthReferenceNumber  string                                              `json:"amexVerbalAuthReferenceNumber,omitempty"`
	MasterCardServiceCode          string                                              `json:"masterCardServiceCode,omitempty"`
	MasterCardServiceReplyCode     string                                              `json:"masterCardServiceReplyCode,omitempty"`
	MasterCardAuthenticationType   string                                              `json:"masterCardAuthenticationType,omitempty"`
	Name                           string                                              `json:"name,omitempty"`
	MerchantNumber                 string                                              `json:"merchantNumber,omitempty"`
	RetrievalReferenceNumber       string                                              `json:"retrievalReferenceNumber,omitempty"`
	PaymentUrl                     string                                              `json:"paymentUrl,omitempty"`
	CompleteUrl                    string                                              `json:"completeUrl,omitempty"`
	Signature                      string                                              `json:"signature,omitempty"`
	PublicKey                      string                                              `json:"publicKey,omitempty"`
	AVS                            *ProcessorInformationAvs                            `json:"avs,omitempty"`
	CardVerification               *ProcessorInformationCardVerification               `json:"cardVerification,omitempty"`
	MerchantAdvice                 *ProcessorInformationMerchantAdvice                 `json:"merchantAdvice,omitempty"`
	ElectronicVerificationResults  *ProcessorInformationElectronicVerificationResults  `json:"electronicVerificationResults,omitempty"`
	AchVerification                *ProcessorInformationAchVerification                `json:"achVerification,omitempty"`
	Customer                       *ProcessorInformationCustomer                       `json:"token_management,omitempty"`
	ConsumerAuthenticationResponse *ProcessorInformationConsumerAuthenticationResponse `json:"consumerAuthenticationResponse,omitempty"`
	Routing                        *ProcessorInformationRouting                        `json:"routing,omitempty"`
}

type ProcessorInformationAvs struct {
	Code    string `json:"code,omitempty"`
	CodeRaw string `json:"codeRaw,omitempty"`
}

type ProcessorInformationCardVerification struct {
	ResultCode    string `json:"resultCode,omitempty"`
	ResultCodeRaw string `json:"resultCodeRaw,omitempty"`
}

type ProcessorInformationMerchantAdvice struct {
	Code      string `json:"code,omitempty"`
	CodeRaw   string `json:"codeRaw,omitempty"`
	NameMatch string `json:"nameMatch,omitempty"`
}

type ProcessorInformationElectronicVerificationResults struct {
	Code           string `json:"code,omitempty"`
	CodeRaw        string `json:"codeRaw,omitempty"`
	Email          string `json:"email,omitempty"`
	EmailRaw       string `json:"emailRaw,omitempty"`
	PhoneNumber    string `json:"phoneNumber,omitempty"`
	PhoneNumberRaw string `json:"phoneNumberRaw,omitempty"`
	PostalCode     string `json:"postalCode,omitempty"`
	PostalCodeRaw  string `json:"postalCodeRaw,omitempty"`
	Street         string `json:"street,omitempty"`
	StreetRaw      string `json:"streetRaw,omitempty"`
	Name           string `json:"name,omitempty"`
	NameRaw        string `json:"nameRaw,omitempty"`
	FirstName      string `json:"firstName,omitempty"`
	FirstNameRaw   string `json:"firstNameRaw,omitempty"`
	MiddleName     string `json:"middleName,omitempty"`
	MiddleNameRaw  string `json:"middleNameRaw,omitempty"`
	LastName       string `json:"lastName,omitempty"`
	LastNameRaw    string `json:"lastNameRaw,omitempty"`
}

type ProcessorInformationAchVerification struct {
	ResultCode    string `json:"resultCode,omitempty"`
	ResultCodeRaw string `json:"resultCodeRaw,omitempty"`
}

type ProcessorInformationCustomer struct {
	PersonalIDResult string `json:"personalIdResult,omitempty"`
}

type ProcessorInformationConsumerAuthenticationResponse struct {
	Code    string `json:"code,omitempty"`
	CodeRaw string `json:"codeRaw,omitempty"`
}

type ProcessorInformationRouting struct {
	Network                   string `json:"network,omitempty"`
	NetworkName               string `json:"networkName,omitempty"`
	CustomerSignatureRequired string `json:"customerSignatureRequired,omitempty"`
}
