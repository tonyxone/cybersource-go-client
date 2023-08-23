package request

type ProcessingInformation struct {
	ActionList                 []string                                         `json:"actionList,omitempty"`
	EnableEscrowOption         bool                                             `json:"enableEscrowOption,omitempty"`
	ActionTokenTypes           []string                                         `json:"actionTokenTypes,omitempty"`
	BinSource                  string                                           `json:"binSource,omitempty"`
	Capture                    bool                                             `json:"capture,omitempty"`
	ProcessorID                string                                           `json:"processorId,omitempty"`
	BusinessApplicationID      string                                           `json:"businessApplicationId,omitempty"`
	CommerceIndicator          string                                           `json:"commerceIndicator,omitempty"`
	CommerceIndicatorLabel     string                                           `json:"commerceIndicatorLabel,omitempty"`
	PaymentSolution            string                                           `json:"paymentSolution,omitempty"`
	ReconciliationID           string                                           `json:"reconciliationId,omitempty"`
	LinkID                     string                                           `json:"linkId,omitempty"`
	PurchaseLevel              string                                           `json:"purchaseLevel,omitempty"`
	PaymentID                  string                                           `json:"paymentId,omitempty"`
	ReportGroup                string                                           `json:"reportGroup,omitempty"`
	VisaCheckoutID             string                                           `json:"visaCheckoutId,omitempty"`
	IndustryDataType           string                                           `json:"industryDataType,omitempty"`
	AuthorizationOptions       *ProcessingInformationAuthorizationOptions       `json:"authorizationOptions,omitempty"`
	CaptureOptions             *ProcessingInformationCaptureOptions             `json:"captureOptions,omitempty"`
	RecurringOptions           *ProcessingInformationRecurringOptions           `json:"recurringOptions,omitempty"`
	BankTransferOptions        *ProcessingInformationBankTransferOptions        `json:"bankTransferOptions,omitempty"`
	PurchaseOptions            *ProcessingInformationPurchaseOptions            `json:"purchaseOptions,omitempty"`
	ElectronicBenefitsTransfer *ProcessingInformationElectronicBenefitsTransfer `json:"electronicBenefitsTransfer,omitempty"`
	LoanOptions                *ProcessingInformationLoanOptions                `json:"loanOptions,omitempty"`
	WalletType                 string                                           `json:"walletType,omitempty"`
	NationalNetDomesticData    string                                           `json:"nationalNetDomesticData,omitempty"`
	JapanPaymentOptions        *ProcessingInformationJapanPaymentOptions        `json:"japanPaymentOptions,omitempty"`
	MobileRemotePaymentType    string                                           `json:"mobileRemotePaymentType,omitempty"`
	ExtendedCreditTotalCount   string                                           `json:"extendedCreditTotalCount,omitempty"`
	NetworkRoutingOrder        string                                           `json:"networkRoutingOrder,omitempty"`
	PayByPointsIndicator       bool                                             `json:"payByPointsIndicator,omitempty"`
	IsReturnAuthRecordEnabled  bool                                             `json:"isReturnAuthRecordEnabled,omitempty"`
}

type ProcessingInformationAuthorizationOptions struct {
	AuthType                string                                              `json:"authType,omitempty"`
	PanReturnIndicator      string                                              `json:"panReturnIndicator,omitempty"`
	VerbalAuthCode          string                                              `json:"verbalAuthCode,omitempty"`
	VerbalAuthTransactionID string                                              `json:"verbalAuthTransactionId,omitempty"`
	AuthIndicator           string                                              `json:"authIndicator,omitempty"`
	PartialAuthIndicator    bool                                                `json:"partialAuthIndicator,omitempty"`
	BalanceInquiry          bool                                                `json:"balanceInquiry,omitempty"`
	IgnoreAvsResult         bool                                                `json:"ignoreAvsResult,omitempty"`
	DeclineAvsFlags         []string                                            `json:"declineAvsFlags,omitempty"`
	IgnoreCvResult          bool                                                `json:"ignoreCvResult,omitempty"`
	Initiator               *ProcessingInformationAuthorizationOptionsInitiator `json:"initiator,omitempty"`
	BillPayment             bool                                                `json:"billPayment,omitempty"`
	BillPaymentType         string                                              `json:"billPaymentType,omitempty"`
	RedemptionInquiry       bool                                                `json:"redemptionInquiry,omitempty"`
	TransportationMode      string                                              `json:"transportationMode,omitempty"`
	AggregatedAuthIndicator string                                              `json:"aggregatedAuthIndicator,omitempty"`
	DebtRecoveryIndicator   string                                              `json:"debtRecoveryIndicator,omitempty"`
	DeferredAuthIndicator   bool                                                `json:"deferredAuthIndicator,omitempty"`
	CashAdvanceIndicator    bool                                                `json:"cashAdvanceIndicator,omitempty"`
	SplitPaymentTransaction bool
}

type ProcessingInformationCaptureOptions struct {
	CaptureSequenceNumber int    `json:"captureSequenceNumber,omitempty"`
	TotalCaptureCount     int    `json:"totalCaptureCount,omitempty"`
	DateToCapture         string `json:"dateToCapture,omitempty"`
}

type ProcessingInformationRecurringOptions struct {
	LoanPayment           bool `json:"loanPayment,omitempty"`
	FirstRecurringPayment bool `json:"firstRecurringPayment,omitempty"`
}

type ProcessingInformationBankTransferOptions struct {
	DeclineAvsFlags     string `json:"declineAvsFlags,omitempty"`
	SecCode             string `json:"secCode,omitempty"`
	TerminalCity        string `json:"terminalCity,omitempty"`
	TerminalState       string `json:"terminalState,omitempty"`
	EffectiveDate       string `json:"effectiveDate,omitempty"`
	PartialPaymentID    string `json:"partialPaymentId,omitempty"`
	CustomerMemo        string `json:"customerMemo,omitempty"`
	PaymentCategoryCode string `json:"paymentCategoryCode,omitempty"`
	SettlementMethod    string `json:"settlementMethod,omitempty"`
	FraudScreeningLevel string `json:"fraudScreeningLevel,omitempty"`
	CustomerPresent     string `json:"customerPresent,omitempty"`
}

type ProcessingInformationPurchaseOptions struct {
	IsElectronicBenefitsTransfer bool   `json:"isElectronicBenefitsTransfer,omitempty"`
	Type                         string `json:"type,omitempty"`
}

type ProcessingInformationElectronicBenefitsTransfer struct {
	Category            string `json:"category,omitempty"`
	VoucherSerialNumber string `json:"voucherSerialNumber,omitempty"`
}

type ProcessingInformationLoanOptions struct {
	Type      string `json:"type,omitempty"`
	AssetType string `json:"assetType,omitempty"`
}

type ProcessingInformationJapanPaymentOptions struct {
	PaymentMethod            string `json:"paymentMethod,omitempty"`
	Bonuses                  string `json:"bonuses,omitempty"`
	BonusMonth               string `json:"bonusMonth,omitempty"`
	SecondBonusMonth         string `json:"secondBonusMonth,omitempty"`
	BonusAmount              string `json:"bonusAmount,omitempty"`
	SecondBonusAmount        string `json:"secondBonusAmount,omitempty"`
	PreapprovalType          string `json:"preapprovalType,omitempty"`
	Installments             string `json:"installments,omitempty"`
	TerminalId               string `json:"terminalId,omitempty"`
	FirstBillingMonth        string `json:"firstBillingMonth,omitempty"`
	BusinessName             string `json:"businessName,omitempty"`
	BusinessNameKatakana     string `json:"businessNameKatakana,omitempty"`
	Jis2TrackData            string `json:"jis2TrackData,omitempty"`
	BusinessNameAlphaNumeric string `json:"businessNameAlphaNumeric,omitempty"`
}

type ProcessingInformationAuthorizationOptionsInitiator struct {
	Type                         string                                                                          `json:"type,omitempty"`
	CredentialStoredOnFile       bool                                                                            `json:"credentialStoredOnFile,omitempty"`
	StoredCredentialUsed         bool                                                                            `json:"storedCredentialUsed,omitempty"`
	MerchantInitiatedTransaction *ProcessingInformationAuthorizationOptionsInitiatorMerchantInitiatedTransaction `json:"merchantInitiatedTransaction,omitempty"`
}

type ProcessingInformationAuthorizationOptionsInitiatorMerchantInitiatedTransaction struct {
	Reason                   string `json:"reason,omitempty"`
	PreviousTransactionID    string `json:"previousTransactionId,omitempty"`
	OriginalAuthorizedAmount string `json:"originalAuthorizedAmount,omitempty"`
}
