package response

type InstrumentIdentifierResponse struct {
	ID     string `json:"id"`
	Object string `json:"object"`
	State  string `json:"state"`
	Type   string `json:"type"`
}
