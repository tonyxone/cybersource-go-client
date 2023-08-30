package response

type WatchlistScreeningInformation struct {
	IpCountryConfidence int                                     `json:"ipCountryConfidence,omitempty"`
	InfoCodes           []string                                `json:"infoCodes,omitempty"`
	WatchList           *WatchlistScreeningInformationWatchList `json:"watchList,omitempty"`
}

type WatchlistScreeningInformationWatchList struct {
	Matches []*WatchlistScreeningInformationWatchListMatches `json:"matches,omitempty"`
}

type WatchlistScreeningInformationWatchListMatches struct {
	Addresses    []string `json:"addresses,omitempty"`
	SanctionList string   `json:"sanctionList,omitempty"`
	Aliases      []string `json:"aliases,omitempty"`
	Programs     []string `json:"programs,omitempty"`
}
