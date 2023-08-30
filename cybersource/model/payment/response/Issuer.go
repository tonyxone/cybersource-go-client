package response

type IssuerInformation struct {
	Country                          string `json:"country,omitempty"`
	DiscretionaryData                string `json:"discretionaryData,omitempty"`
	CountrySpecificDiscretionaryData string `json:"countrySpecificDiscretionaryData,omitempty"`
	ResponseCode                     string `json:"responseCode,omitempty"`
}
