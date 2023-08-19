package model

type DeviceInformation struct {
	HostName                     string                      `json:"hostName,omitempty"`
	IpAddress                    string                      `json:"ipAddress,omitempty"`
	UserAgent                    string                      `json:"userAgent,omitempty"`
	FingerprintSessionId         string                      `json:"fingerprintSessionId,omitempty"`
	UseRawFingerprintSessionId   *bool                       `json:"useRawFingerprintSessionId,omitempty"`
	DeviceType                   string                      `json:"deviceType,omitempty"`
	AppUrl                       string                      `json:"appUrl,omitempty"`
	RawData                      []*DeviceInformationRawData `json:"rawData,omitempty"`
	HttpAcceptBrowserValue       string                      `json:"httpAcceptBrowserValue,omitempty"`
	HttpAcceptContent            string                      `json:"httpAcceptContent,omitempty"`
	HttpBrowserEmail             string                      `json:"httpBrowserEmail,omitempty"`
	HttpBrowserLanguage          string                      `json:"httpBrowserLanguage,omitempty"`
	HttpBrowserJavaEnabled       *bool                       `json:"httpBrowserJavaEnabled,omitempty"`
	HttpBrowserJavaScriptEnabled *bool                       `json:"httpBrowserJavaScriptEnabled,omitempty"`
	HttpBrowserColorDepth        string                      `json:"httpBrowserColorDepth,omitempty"`
	HttpBrowserScreenHeight      string                      `json:"httpBrowserScreenHeight,omitempty"`
	HttpBrowserScreenWidth       string                      `json:"httpBrowserScreenWidth,omitempty"`
	HttpBrowserTimeDifference    string                      `json:"httpBrowserTimeDifference,omitempty"`
	UserAgentBrowserValue        string                      `json:"userAgentBrowserValue,omitempty"`
}

type DeviceInformationRawData struct {
	Data     string `json:"data,omitempty"`
	Provider string `json:"provider,omitempty"`
}
