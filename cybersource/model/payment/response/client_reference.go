package response

type ClientReferenceInformation struct {
	Code                string `json:"code,omitempty"`
	SubmitLocalDateTime string `json:"submitLocalDateTime,omitempty"`
	OwnerMerchantID     string `json:"ownerMerchantId,omitempty"`
}
