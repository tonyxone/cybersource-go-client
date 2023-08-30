package response

type ConsumerAuthenticationInformation struct {
	AccessToken                      string                                                `json:"accessToken,omitempty"`
	AcsRenderingType                 string                                                `json:"acsRenderingType,omitempty"`
	AcsTransactionId                 string                                                `json:"acsTransactionId,omitempty"`
	AcsUrl                           string                                                `json:"acsUrl,omitempty"`
	AuthenticationPath               string                                                `json:"authenticationPath,omitempty"`
	AuthorizationPayload             string                                                `json:"authorizationPayload,omitempty"`
	AuthenticationTransactionId      string                                                `json:"authenticationTransactionId,omitempty"`
	CardholderMessage                string                                                `json:"cardholderMessage,omitempty"`
	Cavv                             string                                                `json:"cavv,omitempty"`
	CavvAlgorithm                    string                                                `json:"cavvAlgorithm,omitempty"`
	ChallengeCancelCode              string                                                `json:"challengeCancelCode,omitempty"`
	ChallengeRequired                string                                                `json:"challengeRequired,omitempty"`
	DecoupledAuthenticationIndicator string                                                `json:"decoupledAuthenticationIndicator,omitempty"`
	DirectoryServerErrorCode         string                                                `json:"directoryServerErrorCode,omitempty"`
	DirectoryServerErrorDescription  string                                                `json:"directoryServerErrorDescription,omitempty"`
	EcommerceIndicator               string                                                `json:"ecommerceIndicator,omitempty"`
	Eci                              string                                                `json:"eci,omitempty"`
	EciRaw                           string                                                `json:"eciRaw,omitempty"`
	EffectiveAuthenticationType      string                                                `json:"effectiveAuthenticationType,omitempty"`
	NetworkScore                     string                                                `json:"networkScore,omitempty"`
	Pareq                            string                                                `json:"pareq,omitempty"`
	ParesStatus                      string                                                `json:"paresStatus,omitempty"`
	ProofXml                         string                                                `json:"proofXml,omitempty"`
	ProxyPan                         string                                                `json:"proxyPan,omitempty"`
	SdkTransactionId                 string                                                `json:"sdkTransactionId,omitempty"`
	SignedParesStatusReason          string                                                `json:"signedParesStatusReason,omitempty"`
	SpecificationVersion             string                                                `json:"specificationVersion,omitempty"`
	StepUpUrl                        string                                                `json:"stepUpUrl,omitempty"`
	ThreeDSServerTransactionId       string                                                `json:"threeDSServerTransactionId,omitempty"`
	UcafAuthenticationData           string                                                `json:"ucafAuthenticationData,omitempty"`
	UcafCollectionIndicator          string                                                `json:"ucafCollectionIndicator,omitempty"`
	VeresEnrolled                    string                                                `json:"veresEnrolled,omitempty"`
	WhiteListStatusSource            string                                                `json:"whiteListStatusSource,omitempty"`
	Xid                              string                                                `json:"xid,omitempty"`
	DirectoryServerTransactionId     string                                                `json:"directoryServerTransactionId,omitempty"`
	AuthenticationResult             string                                                `json:"authenticationResult,omitempty"`
	AuthenticationStatusMsg          string                                                `json:"authenticationStatusMsg,omitempty"`
	Indicator                        string                                                `json:"indicator,omitempty"`
	InteractionCounter               string                                                `json:"interactionCounter,omitempty"`
	WhiteListStatus                  string                                                `json:"whiteListStatus,omitempty"`
	Ivr                              ConsumerAuthenticationInformationIvr                  `json:"ivr,omitempty"`
	StrongAuthentication             ConsumerAuthenticationInformationStrongAuthentication `json:"strongAuthentication,omitempty"`
}

type ConsumerAuthenticationInformationIvr struct {
	EnabledMessage      bool   `json:"enabledMessage,omitempty"`
	EncryptionKey       string `json:"encryptionKey,omitempty"`
	EncryptionMandatory bool   `json:"encryptionMandatory,omitempty"`
	EncryptionType      string `json:"encryptionType,omitempty"`
	Label               string `json:"label,omitempty"`
	Prompt              string `json:"prompt,omitempty"`
	StatusMessage       string `json:"statusMessage,omitempty"`
}

type ConsumerAuthenticationInformationStrongAuthentication struct {
	IssuerInformation ConsumerAuthenticationInformationStrongAuthenticationIssuerInformation `json:"issuerInformation,omitempty"`
}

type ConsumerAuthenticationInformationStrongAuthenticationIssuerInformation struct {
	RiskAnalysisExemptionResult            string `json:"riskAnalysisExemptionResult,omitempty"`
	TrustedMerchantExemptionResult         string `json:"trustedMerchantExemptionResult,omitempty"`
	LowValueExemptionResult                string `json:"lowValueExemptionResult,omitempty"`
	SecureCorporatePaymentResult           string `json:"secureCorporatePaymentResult,omitempty"`
	TransactionRiskAnalysisExemptionResult string `json:"transactionRiskAnalysisExemptionResult,omitempty"`
}
