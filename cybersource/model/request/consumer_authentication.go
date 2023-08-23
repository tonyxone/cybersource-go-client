package request

type ConsumerAuthenticationInformation struct {
	Cavv                                      string                                                 `json:"cavv,omitempty"`
	CavvAlgorithm                             string                                                 `json:"cavvAlgorithm,omitempty"`
	EciRaw                                    string                                                 `json:"eciRaw,omitempty"`
	ParesStatus                               string                                                 `json:"paresStatus,omitempty"`
	VeresEnrolled                             string                                                 `json:"veresEnrolled,omitempty"`
	Xid                                       string                                                 `json:"xid,omitempty"`
	UcafCollectionIndicator                   string                                                 `json:"ucafCollectionIndicator,omitempty"`
	UcafAuthenticationData                    string                                                 `json:"ucafAuthenticationData,omitempty"`
	StrongAuthentication                      *ConsumerAuthenticationInformationStrongAuthentication `json:"strongAuthentication,omitempty"`
	DirectoryServerTransactionId              string                                                 `json:"directoryServerTransactionId,omitempty"`
	PaSpecificationVersion                    string                                                 `json:"paSpecificationVersion,omitempty"`
	AuthenticationType                        string                                                 `json:"authenticationType,omitempty"`
	ResponseAccessToken                       string                                                 `json:"responseAccessToken,omitempty"`
	AcsTransactionId                          string                                                 `json:"acsTransactionId,omitempty"`
	AcsWindowSize                             string                                                 `json:"acsWindowSize,omitempty"`
	AlternateAuthenticationData               string                                                 `json:"alternateAuthenticationData,omitempty"`
	AlternateAuthenticationDate               string                                                 `json:"alternateAuthenticationDate,omitempty"`
	AlternateAuthenticationMethod             string                                                 `json:"alternateAuthenticationMethod,omitempty"`
	AuthenticationDate                        string                                                 `json:"authenticationDate,omitempty"`
	AuthenticationTransactionId               string                                                 `json:"authenticationTransactionId,omitempty"`
	ChallengeCancelCode                       string                                                 `json:"challengeCancelCode,omitempty"`
	ChallengeCode                             string                                                 `json:"challengeCode,omitempty"`
	ChallengeStatus                           string                                                 `json:"challengeStatus,omitempty"`
	CustomerCardAlias                         string                                                 `json:"customerCardAlias,omitempty"`
	DecoupledAuthenticationIndicator          string                                                 `json:"decoupledAuthenticationIndicator,omitempty"`
	DecoupledAuthenticationMaxTime            string                                                 `json:"decoupledAuthenticationMaxTime,omitempty"`
	DefaultCard                               bool                                                   `json:"defaultCard,omitempty"`
	DeviceChannel                             string                                                 `json:"deviceChannel,omitempty"`
	InstallmentTotalCount                     int                                                    `json:"installmentTotalCount,omitempty"`
	MerchantFraudRate                         string                                                 `json:"merchantFraudRate,omitempty"`
	MarketingOptIn                            bool                                                   `json:"marketingOptIn,omitempty"`
	MarketingSource                           string                                                 `json:"marketingSource,omitempty"`
	Mcc                                       string                                                 `json:"mcc,omitempty"`
	MerchantScore                             int                                                    `json:"merchantScore,omitempty"`
	MessageCategory                           string                                                 `json:"messageCategory,omitempty"`
	NetworkScore                              string                                                 `json:"networkScore,omitempty"`
	NpaCode                                   string                                                 `json:"npaCode,omitempty"`
	OverridePaymentMethod                     string                                                 `json:"overridePaymentMethod,omitempty"`
	OverrideCountryCode                       string                                                 `json:"overrideCountryCode,omitempty"`
	PriorAuthenticationData                   string                                                 `json:"priorAuthenticationData,omitempty"`
	PriorAuthenticationMethod                 string                                                 `json:"priorAuthenticationMethod,omitempty"`
	PriorAuthenticationReferenceId            string                                                 `json:"priorAuthenticationReferenceId,omitempty"`
	PriorAuthenticationTime                   string                                                 `json:"priorAuthenticationTime,omitempty"`
	ProductCode                               string                                                 `json:"productCode,omitempty"`
	ReturnUrl                                 string                                                 `json:"returnUrl,omitempty"`
	RequestorId                               string                                                 `json:"requestorId,omitempty"`
	RequestorInitiatedAuthenticationIndicator string                                                 `json:"requestorInitiatedAuthenticationIndicator,omitempty"`
	RequestorName                             string                                                 `json:"requestorName,omitempty"`
	ReferenceId                               string                                                 `json:"referenceId,omitempty"`
	SdkMaxTimeout                             string                                                 `json:"sdkMaxTimeout,omitempty"`
	SecureCorporatePaymentIndicator           string                                                 `json:"secureCorporatePaymentIndicator,omitempty"`
	TransactionMode                           string                                                 `json:"transactionMode,omitempty"`
	WhiteListStatus                           string                                                 `json:"whiteListStatus,omitempty"`
	EffectiveAuthenticationType               string                                                 `json:"effectiveAuthenticationType,omitempty"`
	SignedParesStatusReason                   string                                                 `json:"signedParesStatusReason,omitempty"`
	SignedPares                               string                                                 `json:"signedPares,omitempty"`
}

type ConsumerAuthenticationInformationStrongAuthentication struct {
	LowValueExemptionIndicator                string `json:"lowValueExemptionIndicator,omitempty"`
	RiskAnalysisExemptionIndicator            string `json:"riskAnalysisExemptionIndicator,omitempty"`
	TrustedMerchantExemptionIndicator         string `json:"trustedMerchantExemptionIndicator,omitempty"`
	SecureCorporatePaymentIndicator           string `json:"secureCorporatePaymentIndicator,omitempty"`
	DelegatedAuthenticationExemptionIndicator string `json:"delegatedAuthenticationExemptionIndicator,omitempty"`
	OutageExemptionIndicator                  string `json:"outageExemptionIndicator,omitempty"`
	AuthenticationIndicator                   string `json:"authenticationIndicator,omitempty"`
}
