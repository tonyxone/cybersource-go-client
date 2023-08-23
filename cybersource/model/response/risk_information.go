package response

type RiskInformation struct {
	CasePriority int                          `json:"casePriority,omitempty"`
	LocalTime    string                       `json:"localTime,omitempty"`
	Providers    map[string]map[string]string `json:"providers,omitempty"`
	Profile      *RiskInformationProfile      `json:"profile,omitempty"`
	Rules        []*RiskInformationRules      `json:"rules,omitempty"`
	InfoCodes    *RiskInformationInfoCodes    `json:"infoCodes,omitempty"`
	Velocity     *RiskInformationVelocity     `json:"velocity,omitempty"`
	Score        *RiskInformationScore        `json:"score,omitempty"`
	IPAddress    *RiskInformationIPAddress    `json:"ipAddress,omitempty"`
	Travel       *RiskInformationTravel       `json:"travel,omitempty"`
}

type RiskInformationProfile struct {
	Name             string `json:"name,omitempty"`
	DestinationQueue string `json:"destinationQueue,omitempty"`
	SelectorRule     string `json:"selectorRule,omitempty"`
}

type RiskInformationRules struct {
	Name     string `json:"name,omitempty"`
	Decision string `json:"decision,omitempty"`
}

type RiskInformationInfoCodes struct {
	Velocity       []string `json:"velocity,omitempty"`
	Address        []string `json:"address,omitempty"`
	CustomerList   []string `json:"customerList,omitempty"`
	DeviceBehavior []string `json:"deviceBehavior,omitempty"`
	IdentityChange []string `json:"identityChange,omitempty"`
	Internet       []string `json:"internet,omitempty"`
	Phone          []string `json:"phone,omitempty"`
	Suspicious     []string `json:"suspicious,omitempty"`
	GlobalVelocity []string `json:"globalVelocity,omitempty"`
}

type RiskInformationVelocity struct {
	Morphing []*RiskInformationVelocityMorphing `json:"morphing,omitempty"`
	Address  []string                           `json:"address,omitempty"`
}

type RiskInformationVelocityMorphing struct {
	Count           int    `json:"count,omitempty"`
	FieldName       string `json:"fieldName,omitempty"`
	InformationCode string `json:"informationCode,omitempty"`
}

type RiskInformationScore struct {
	FactorCodes []string `json:"factorCodes,omitempty"`
	ModelUsed   string   `json:"modelUsed,omitempty"`
	Result      string   `json:"result,omitempty"`
}

type RiskInformationIPAddress struct {
	AnonymizerStatus   string `json:"anonymizerStatus,omitempty"`
	Locality           string `json:"locality,omitempty"`
	Country            string `json:"country,omitempty"`
	AdministrativeArea string `json:"administrativeArea,omitempty"`
	RoutingMethod      string `json:"routingMethod,omitempty"`
	Carrier            string `json:"carrier,omitempty"`
	Organization       string `json:"organization,omitempty"`
}

type RiskInformationTravel struct {
	ActualFinalDestination *RiskInformationTravelActualFinalDestination `json:"actualFinalDestination,omitempty"`
	FirstDeparture         *RiskInformationTravelFirstDeparture         `json:"firstDeparture,omitempty"`
	FirstDestination       *RiskInformationTravelFirstDestination       `json:"firstDestination,omitempty"`
	LastDestination        *RiskInformationTravelLastDestination        `json:"lastDestination,omitempty"`
}

type RiskInformationTravelActualFinalDestination struct {
	Country   string `json:"country,omitempty"`
	Locality  string `json:"locality,omitempty"`
	Latitude  string `json:"latitude,omitempty"`
	Longitude string `json:"longitude,omitempty"`
}

type RiskInformationTravelFirstDeparture struct {
	Country   string `json:"country,omitempty"`
	Locality  string `json:"locality,omitempty"`
	Latitude  string `json:"latitude,omitempty"`
	Longitude string `json:"longitude,omitempty"`
}

type RiskInformationTravelFirstDestination struct {
	Country   string `json:"country,omitempty"`
	Locality  string `json:"locality,omitempty"`
	Latitude  string `json:"latitude,omitempty"`
	Longitude string `json:"longitude,omitempty"`
}

type RiskInformationTravelLastDestination struct {
	Country   string `json:"country,omitempty"`
	Locality  string `json:"locality,omitempty"`
	Latitude  string `json:"latitude,omitempty"`
	Longitude string `json:"longitude,omitempty"`
}
