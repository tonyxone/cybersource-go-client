package model

type PointOfSaleInformation struct {
	TerminalId                                   string                     `json:"terminalId,omitempty"`
	TerminalSerialNumber                         string                     `json:"terminalSerialNumber,omitempty"`
	CardholderVerificationMethodUsed             int                        `json:"cardholderVerificationMethodUsed,omitempty"`
	LaneNumber                                   string                     `json:"laneNumber,omitempty"`
	CatLevel                                     int                        `json:"catLevel,omitempty"`
	EntryMode                                    string                     `json:"entryMode,omitempty"`
	TerminalCapability                           int                        `json:"terminalCapability,omitempty"`
	OperatingEnvironment                         string                     `json:"operatingEnvironment,omitempty"`
	Emv                                          *PointOfSaleInformationEmv `json:"emv,omitempty"`
	AmexCapnData                                 string                     `json:"amexCapnData,omitempty"`
	TrackData                                    string                     `json:"trackData,omitempty"`
	StoreAndForwardIndicator                     string                     `json:"storeAndForwardIndicator,omitempty"`
	CardholderVerificationMethod                 []string                   `json:"cardholderVerificationMethod,omitempty"`
	TerminalInputCapability                      []string                   `json:"terminalInputCapability,omitempty"`
	TerminalCardCaptureCapability                string                     `json:"terminalCardCaptureCapability,omitempty"`
	TerminalOutputCapability                     string                     `json:"terminalOutputCapability,omitempty"`
	TerminalPinCapability                        int                        `json:"terminalPinCapability,omitempty"`
	PinEntrySolution                             string                     `json:"pinEntrySolution,omitempty"`
	DeviceId                                     string                     `json:"deviceId,omitempty"`
	PinBlockEncodingFormat                       int                        `json:"pinBlockEncodingFormat,omitempty"`
	EncryptedPin                                 string                     `json:"encryptedPin,omitempty"`
	EncryptedKeySerialNumber                     string                     `json:"encryptedKeySerialNumber,omitempty"`
	PartnerSdkVersion                            string                     `json:"partnerSdkVersion,omitempty"`
	EmvApplicationIdentifierAndDedicatedFileName string                     `json:"emvApplicationIdentifierAndDedicatedFileName,omitempty"`
	TerminalCompliance                           string                     `json:"terminalCompliance,omitempty"`
	IsDedicatedHardwareTerminal                  string                     `json:"isDedicatedHardwareTerminal,omitempty"`
	TerminalModel                                string                     `json:"terminalModel,omitempty"`
	TerminalMake                                 string                     `json:"terminalMake,omitempty"`
	ServiceCode                                  string                     `json:"serviceCode,omitempty"`
}

type PointOfSaleInformationEmv struct {
	Tags                             string `json:"tags,omitempty"`
	CardholderVerificationMethodUsed int    `json:"cardholderVerificationMethodUsed,omitempty"`
	CardSequenceNumber               string `json:"cardSequenceNumber,omitempty"`
	Fallback                         bool   `json:"fallback,omitempty"`
	FallbackCondition                int    `json:"fallbackCondition,omitempty"`
	IsRepeat                         bool   `json:"isRepeat,omitempty"`
}
