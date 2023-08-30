package request

type AcquirerInformation struct {
	AcquirerBin string `json:"acquirerBin,omitempty"`
	Country     string `json:"country,omitempty"`
	Password    string `json:"password,omitempty"`
	MerchantId  string `json:"merchantId,omitempty"`
}
