package response

type PointOfSaleInformation struct {
	Emv          PointOfSaleInformationEmv `json:"emv,omitempty"`
	AmexCapnData string                    `json:"amexCapnData,omitempty"`
	TerminalId   string                    `json:"terminalId,omitempty"`
}

type PointOfSaleInformationEmv struct {
	Tags                 string `json:"tags,omitempty"`
	ChipValidationType   string `json:"chipValidationType,omitempty"`
	ChipValidationResult string `json:"chipValidationResult,omitempty"`
}
