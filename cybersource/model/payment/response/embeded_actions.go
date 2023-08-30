package response

type EmbeddedActions struct {
	Capture                        *EmbeddedActionsCAPTURE                `json:"CAPTURE,omitempty"`
	Decision                       *EmbeddedActionsDECISION               `json:"DECISION,omitempty"`
	ConsumerAuthentication         *EmbeddedActionsCONSUMERAUTHENTICATION `json:"CONSUMER_AUTHENTICATION,omitempty"`
	ValidateConsumerAuthentication *EmbeddedActionsCONSUMERAUTHENTICATION `json:"VALIDATE_CONSUMER_AUTHENTICATION,omitempty"`
	Watchlist_Screening            *EmbeddedActionsWATCHLISTSCREENING     `json:"WATCHLIST_SCREENING,omitempty"`
}

type EmbeddedActionsCAPTURE struct {
	Status  string `json:"status,omitempty"`
	Reason  string `json:"reason,omitempty"`
	Message string `json:"message,omitempty"`
}

type EmbeddedActionsDECISION struct {
	Status  string `json:"status,omitempty"`
	Reason  string `json:"reason,omitempty"`
	Message string `json:"message,omitempty"`
}

type EmbeddedActionsCONSUMERAUTHENTICATION struct {
	Status  string `json:"status,omitempty"`
	Reason  string `json:"reason,omitempty"`
	Message string `json:"message,omitempty"`
}

type EmbeddedActionsWATCHLISTSCREENING struct {
	Status  string `json:"status,omitempty"`
	Reason  string `json:"reason,omitempty"`
	Message string `json:"message,omitempty"`
}
