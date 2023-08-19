package model

type InstallmentInformation struct {
	Amount                  string `json:"amount,omitempty"`
	Frequency               string `json:"frequency,omitempty"`
	PlanType                string `json:"planType,omitempty"`
	TotalAmount             string `json:"totalAmount,omitempty"`
	FirstInstallmentDate    string `json:"firstInstallmentDate,omitempty"`
	InvoiceData             string `json:"invoiceData,omitempty"`
	PaymentType             string `json:"paymentType,omitempty"`
	EligibilityInquiry      string `json:"eligibilityInquiry,omitempty"`
	GracePeriodDuration     string `json:"gracePeriodDuration,omitempty"`
	GracePeriodDurationType string `json:"gracePeriodDurationType,omitempty"`
	FirstInstallmentAmount  string `json:"firstInstallmentAmount,omitempty"`
	ValidationIndicator     string `json:"validationIndicator,omitempty"`
	Identifier              string `json:"identifier,omitempty"`
	Sequence                int    `json:"sequence,omitempty"`
	TotalCount              int    `json:"totalCount,omitempty"`
}
