package response

type GenerateCaptureContextResponse struct {
	CaptureContext string     `json:"captureContext,omitempty"`
	CorrelationID  string     `json:"correlationId,omitempty"`
	Details        []*Details `json:"details,omitempty"`
	Message        string     `json:"message,omitempty"`
	Errors         []*Errors  `json:"errors,omitempty"`
}

type Details struct {
	Message  string `json:"message,omitempty"`
	Location string `json:"location,omitempty"`
}
