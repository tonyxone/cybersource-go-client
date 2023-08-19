package model

type ProcessorInformation struct {
	PreApprovalToken     string                                    `json:"preApprovalToken,omitempty"`
	AuthorizationOptions *ProcessorInformationAuthorizationOptions `json:"authorizationOptions,omitempty"`
}

type ProcessorInformationAuthorizationOptions struct {
	PanReturnIndicator string `json:"panReturnIndicator,omitempty"`
}
