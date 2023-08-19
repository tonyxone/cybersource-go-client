package model

type WatchlistScreeningInformation struct {
	AddressOperator string                     `json:"addressOperator,omitempty"`
	Weights         *WatchlistScreeningWeights `json:"weights,omitempty"`
	SanctionLists   []string                   `json:"sanctionLists,omitempty"`
	ProceedOnMatch  bool                       `json:"proceedOnMatch,omitempty"`
}

type WatchlistScreeningWeights struct {
	Address string `json:"address,omitempty"`
	Company string `json:"company,omitempty"`
	Name    string `json:"name,omitempty"`
}
