package request

type RiskInformation struct {
	Profile       *RiskInformationProfile         `json:"profile,omitempty"`
	EventType     string                          `json:"eventType,omitempty"`
	BuyerHistory  *RiskInformationBuyerHistory    `json:"buyerHistory,omitempty"`
	AuxiliaryData []*RiskInformationAuxiliaryData `json:"auxiliaryData,omitempty"`
}

type RiskInformationProfile struct {
	Name string `json:"name,omitempty"`
}

type RiskInformationBuyerHistory struct {
	CustomerAccount         *RiskInformationBuyerHistoryCustomerAccount `json:"customerAccount,omitempty"`
	AccountHistory          *RiskInformationBuyerHistoryAccountHistory  `json:"accountHistory,omitempty"`
	AccountPurchases        int                                         `json:"accountPurchases,omitempty"`
	AddCardAttempts         int                                         `json:"addCardAttempts,omitempty"`
	PriorSuspiciousActivity bool                                        `json:"priorSuspiciousActivity,omitempty"`
	PaymentAccountHistory   string                                      `json:"paymentAccountHistory,omitempty"`
	PaymentAccountDate      int                                         `json:"paymentAccountDate,omitempty"`
	TransactionCountDay     int                                         `json:"transactionCountDay,omitempty"`
	TransactionCountYear    int                                         `json:"transactionCountYear,omitempty"`
}

type RiskInformationBuyerHistoryCustomerAccount struct {
	LastChangeDate      string `json:"lastChangeDate,omitempty"`
	CreationHistory     string `json:"creationHistory,omitempty"`
	ModificationHistory string `json:"modificationHistory,omitempty"`
	PasswordHistory     string `json:"passwordHistory,omitempty"`
	CreateDate          string `json:"createDate,omitempty"`
	PasswordChangeDate  string `json:"passwordChangeDate,omitempty"`
}

type RiskInformationBuyerHistoryAccountHistory struct {
	FirstUseOfShippingAddress *bool  `json:"firstUseOfShippingAddress,omitempty"`
	ShippingAddressUsageDate  string `json:"shippingAddressUsageDate,omitempty"`
}

type RiskInformationAuxiliaryData struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}
