package request

type ClientReferenceInformation struct {
	Code               string                             `json:"code,omitempty"`
	ReconciliationID   string                             `json:"reconciliationId,omitempty"`
	PausedRequestID    string                             `json:"pausedRequestId,omitempty"`
	TransactionID      string                             `json:"transactionId,omitempty"`
	Comments           string                             `json:"comments,omitempty"`
	Partner            *ClientReferenceInformationPartner `json:"partner,omitempty"`
	ApplicationName    string                             `json:"applicationName,omitempty"`
	ApplicationVersion string                             `json:"applicationVersion,omitempty"`
	ApplicationUser    string                             `json:"applicationUser,omitempty"`
}

type ClientReferenceInformationPartner struct {
	OriginalTransactionID         string `json:"originalTransactionId,omitempty"`
	DeveloperID                   string `json:"developerId,omitempty"`
	SolutionID                    string `json:"solutionId,omitempty"`
	ThirdPartyCertificationNumber string `json:"thirdPartyCertificationNumber,omitempty"`
}
