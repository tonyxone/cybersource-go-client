package response

type CaptureResponse struct {
	ID                         string                      `json:"id,omitempty"`
	SubmitTimeUtc              string                      `json:"submitTimeUtc,omitempty"`
	Status                     string                      `json:"status,omitempty"`
	ReconciliationID           string                      `json:"reconciliationId,omitempty"`
	ClientReferenceInformation *ClientReferenceInformation `json:"clientReferenceInformation,omitempty"`
	ProcessorInformation       *ProcessorInformation       `json:"processorInformation,omitempty"`
	OrderInformation           *OrderInformation           `json:"orderInformation,omitempty"`
	PointOfSaleInformation     *PointOfSaleInformation     `json:"pointOfSaleInformation,omitempty"`
	ProcessingInformation      *ProcessingInformation      `json:"processingInformation,omitempty"`
	// Reason and Message will be populated for 400 and 500 responses
	Reason  string `json:"reason,omitempty"`
	Message string `json:"message,omitempty"`
}
