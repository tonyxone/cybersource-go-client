package request

type GenerateCaptureContextRequest struct {
	TargetOrigins             []string                   `json:"targetOrigins,omitempty"`
	AllowedCardNetworks       []string                   `json:"allowedCardNetworks,omitempty"`
	ClientVersion             string                     `json:"clientVersion,omitempty"`
	CheckoutAPIInitialization *CheckoutApiInitialization `json:"checkoutApiInitialization,omitempty"`
}

type CheckoutApiInitialization struct {
	ProfileID                 string `json:"profile_id,omitempty"`
	AccessKey                 string `json:"access_key,omitempty"`
	ReferenceNumber           string `json:"reference_number,omitempty"`
	TransactionUUID           string `json:"transaction_uuid,omitempty"`
	TransactionType           string `json:"transaction_type,omitempty"`
	Currency                  string `json:"currency,omitempty"`
	Amount                    string `json:"amount,omitempty"`
	Locale                    string `json:"locale,omitempty"`
	OverrideCustomReceiptPage string `json:"override_custom_receipt_page,omitempty"`
	UnsignedFieldNames        string `json:"unsigned_field_names,omitempty"`
}
