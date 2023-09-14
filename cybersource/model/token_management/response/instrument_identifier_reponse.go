package response

type InstrumentIdentifierResponse struct {
	ID     string `json:"id"`
	Object string `json:"object"`
	State  string `json:"state"`
	Type   string `json:"type"`
	// Reason and Message will be populated for 400 and 500 responses
	Reason  string    `json:"reason,omitempty"`
	Message string    `json:"message,omitempty"`
	Errors  []*Errors `json:"errors,omitempty"`
}
