package response

type InstallmentInformation struct {
	AdditionalCosts           string `json:"additionalCosts,omitempty"`
	AdditionalCostsPercentage string `json:"additionalCostsPercentage,omitempty"`
	Amount                    string `json:"amount,omitempty"`
	AmountFunded              string `json:"amountFunded,omitempty"`
	AmountRequestedPercentage string `json:"amountRequestedPercentage,omitempty"`
	AnnualFinancingCost       string `json:"annualFinancingCost,omitempty"`
	AnnualInterestRate        string `json:"annualInterestRate,omitempty"`
	Expenses                  string `json:"expenses,omitempty"`
	ExpensesPercentage        string `json:"expensesPercentage,omitempty"`
	Fees                      string `json:"fees,omitempty"`
	FeesPercentage            string `json:"feesPercentage,omitempty"`
	Frequency                 string `json:"frequency,omitempty"`
	Insurance                 string `json:"insurance,omitempty"`
	InsurancePercentage       string `json:"insurancePercentage,omitempty"`
	InvoiceData               string `json:"invoiceData,omitempty"`
	MonthlyInterestRate       string `json:"monthlyInterestRate,omitempty"`
	PlanType                  string `json:"planType,omitempty"`
	Taxes                     string `json:"taxes,omitempty"`
	TaxesPercentage           string `json:"taxesPercentage,omitempty"`
	TotalAmount               string `json:"totalAmount,omitempty"`
	MinimumTotalCount         string `json:"minimumTotalCount,omitempty"`
	MaximumTotalCount         string `json:"maximumTotalCount,omitempty"`
	FirstInstallmentAmount    string `json:"firstInstallmentAmount,omitempty"`
	FirstInstallmentDate      string `json:"firstInstallmentDate,omitempty"`
	TotalCount                int    `json:"totalCount,omitempty"`
	Sequence                  int    `json:"sequence,omitempty"`
}
