package model

type TravelInformation struct {
	Duration   string                       `json:"duration,omitempty"`
	Agency     *TravelInformationAgency     `json:"agency,omitempty"`
	AutoRental *TravelInformationAutoRental `json:"autoRental,omitempty"`
	Lodging    *TravelInformationLodging    `json:"lodging,omitempty"`
	Transit    *TravelInformationTransit    `json:"transit,omitempty"`
}

type TravelInformationAgency struct {
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}

type TravelInformationAutoRental struct {
	NoShowIndicator             bool                                      `json:"noShowIndicator,omitempty"`
	CustomerName                string                                    `json:"customerName,omitempty"`
	VehicleClass                string                                    `json:"vehicleClass,omitempty"`
	DistanceTravelled           string                                    `json:"distanceTravelled,omitempty"`
	DistanceUnit                string                                    `json:"distanceUnit,omitempty"`
	ReturnDateTime              string                                    `json:"returnDateTime,omitempty"`
	RentalDateTime              string                                    `json:"rentalDateTime,omitempty"`
	MaxFreeDistance             string                                    `json:"maxFreeDistance,omitempty"`
	InsuranceIndicator          bool                                      `json:"insuranceIndicator,omitempty"`
	ProgramCode                 string                                    `json:"programCode,omitempty"`
	ReturnAddress               *TravelInformationAutoRentalReturnAddress `json:"returnAddress,omitempty"`
	RentalAddress               *TravelInformationAutoRentalRentalAddress `json:"rentalAddress,omitempty"`
	AgreementNumber             string                                    `json:"agreementNumber,omitempty"`
	OdometerReading             string                                    `json:"odometerReading,omitempty"`
	VehicleIdentificationNumber string                                    `json:"vehicleIdentificationNumber,omitempty"`
	CompanyId                   string                                    `json:"companyId,omitempty"`
	NumberOfAdditionalDrivers   string                                    `json:"numberOfAdditionalDrivers,omitempty"`
	DriverAge                   string                                    `json:"driverAge,omitempty"`
	SpecialProgramCode          string                                    `json:"specialProgramCode,omitempty"`
	VehicleMake                 string                                    `json:"vehicleMake,omitempty"`
	VehicleModel                string                                    `json:"vehicleModel,omitempty"`
	TimePeriod                  string                                    `json:"timePeriod,omitempty"`
	CommodityCode               string                                    `json:"commodityCode,omitempty"`
	CustomerServicePhoneNumber  string                                    `json:"customerServicePhoneNumber,omitempty"`
	TaxDetails                  *TravelInformationAutoRentalTaxDetails    `json:"taxDetails,omitempty"`
	InsuranceAmount             string                                    `json:"insuranceAmount,omitempty"`
	OneWayDropOffAmount         string                                    `json:"oneWayDropOffAmount,omitempty"`
	AdjustedAmountIndicator     string                                    `json:"adjustedAmountIndicator,omitempty"`
	AdjustedAmount              string                                    `json:"adjustedAmount,omitempty"`
	FuelCharges                 string                                    `json:"fuelCharges,omitempty"`
	WeeklyRentalRate            string                                    `json:"weeklyRentalRate,omitempty"`
	DailyRentalRate             string                                    `json:"dailyRentalRate,omitempty"`
	RatePerMile                 string                                    `json:"ratePerMile,omitempty"`
	MileageCharge               string                                    `json:"mileageCharge,omitempty"`
	ExtraMileageCharge          string                                    `json:"extraMileageCharge,omitempty"`
	LateFeeAmount               string                                    `json:"lateFeeAmount,omitempty"`
	TowingCharge                string                                    `json:"towingCharge,omitempty"`
	ExtraCharge                 string                                    `json:"extraCharge,omitempty"`
	GpsCharge                   string                                    `json:"gpsCharge,omitempty"`
	PhoneCharge                 string                                    `json:"phoneCharge,omitempty"`
	ParkingViolationCharge      string                                    `json:"parkingViolationCharge,omitempty"`
	OtherCharges                string                                    `json:"otherCharges,omitempty"`
}

type TravelInformationAutoRentalReturnAddress struct {
	City       string `json:"city,omitempty"`
	State      string `json:"state,omitempty"`
	Country    string `json:"country,omitempty"`
	LocationId string `json:"locationId,omitempty"`
	Location   string `json:"location,omitempty"`
}

type TravelInformationAutoRentalRentalAddress struct {
	City       string `json:"city,omitempty"`
	State      string `json:"state,omitempty"`
	Country    string `json:"country,omitempty"`
	LocationId string `json:"locationId,omitempty"`
	Address1   string `json:"address1,omitempty"`
	Address2   string `json:"address2,omitempty"`
	Location   string `json:"location,omitempty"`
}

type TravelInformationAutoRentalTaxDetails struct {
	Amount        string `json:"amount,omitempty"`
	Rate          string `json:"rate,omitempty"`
	Applied       bool   `json:"applied,omitempty"`
	ExemptionCode string `json:"exemptionCode,omitempty"`
	TaxType       string `json:"taxType,omitempty"`
	TaxSummary    string `json:"taxSummary,omitempty"`
}

type TravelInformationLodging struct {
	CheckInDate          string                          `json:"checkInDate,omitempty"`
	CheckOutDate         string                          `json:"checkOutDate,omitempty"`
	Room                 []*TravelInformationLodgingRoom `json:"room,omitempty"`
	SmokingPreference    string                          `json:"smokingPreference,omitempty"`
	NumberOfRooms        int                             `json:"numberOfRooms,omitempty"`
	NumberOfGuests       int                             `json:"numberOfGuests,omitempty"`
	RoomBedType          string                          `json:"roomBedType,omitempty"`
	RoomTaxType          string                          `json:"roomTaxType,omitempty"`
	RoomRateType         string                          `json:"roomRateType,omitempty"`
	GuestName            string                          `json:"guestName,omitempty"`
	CustomerServicePhone string                          `json:"customerServicePhoneNumber,omitempty"`
	CorporateClientCode  string                          `json:"corporateClientCode,omitempty"`
	AdditionalDiscount   string                          `json:"additionalDiscountAmount,omitempty"`
	RoomLocation         string                          `json:"roomLocation,omitempty"`
	SpecialProgramCode   string                          `json:"specialProgramCode,omitempty"`
	TotalTaxAmount       string                          `json:"totalTaxAmount,omitempty"`
	PrepaidCost          string                          `json:"prepaidCost,omitempty"`
	FoodAndBeverageCost  string                          `json:"foodAndBeverageCost,omitempty"`
	RoomTaxAmount        string                          `json:"roomTaxAmount,omitempty"`
	AdjustmentAmount     string                          `json:"adjustmentAmount,omitempty"`
	PhoneCost            string                          `json:"phoneCost,omitempty"`
	RestaurantCost       string                          `json:"restaurantCost,omitempty"`
	RoomServiceCost      string                          `json:"roomServiceCost,omitempty"`
	MiniBarCost          string                          `json:"miniBarCost,omitempty"`
	LaundryCost          string                          `json:"laundryCost,omitempty"`
	MiscellaneousCost    string                          `json:"miscellaneousCost,omitempty"`
	GiftShopCost         string                          `json:"giftShopCost,omitempty"`
	MovieCost            string                          `json:"movieCost,omitempty"`
	HealthClubCost       string                          `json:"healthClubCost,omitempty"`
	ValetParkingCost     string                          `json:"valetParkingCost,omitempty"`
	CashDisbursementCost string                          `json:"cashDisbursementCost,omitempty"`
	NonRoomCost          string                          `json:"nonRoomCost,omitempty"`
	BusinessCenterCost   string                          `json:"businessCenterCost,omitempty"`
	LoungeBarCost        string                          `json:"loungeBarCost,omitempty"`
	TransportationCost   string                          `json:"transportationCost,omitempty"`
	GratuityAmount       string                          `json:"gratuityAmount,omitempty"`
	ConferenceRoomCost   string                          `json:"conferenceRoomCost,omitempty"`
	AudioVisualCost      string                          `json:"audioVisualCost,omitempty"`
	BanquetCost          string                          `json:"banquestCost,omitempty"`
	NonRoomTaxAmount     string                          `json:"nonRoomTaxAmount,omitempty"`
	EarlyCheckOutCost    string                          `json:"earlyCheckOutCost,omitempty"`
	InternetAccessCost   string                          `json:"internetAccessCost,omitempty"`
}

type TravelInformationLodgingRoom struct {
	DailyRate      string `json:"dailyRate,omitempty"`
	NumberOfNights int    `json:"numberOfNights,omitempty"`
}

type TravelInformationTransit struct {
	Airline *TravelInformationTransitAirline `json:"airline,omitempty"`
}

type TravelInformationTransitAirline struct {
	BookingReferenceNumber     string                                               `json:"bookingReferenceNumber,omitempty"`
	CarrierName                string                                               `json:"carrierName,omitempty"`
	TicketIssuer               *TravelInformationTransitAirlineTicketIssuer         `json:"ticketIssuer,omitempty"`
	TicketNumber               string                                               `json:"ticketNumber,omitempty"`
	CheckDigit                 string                                               `json:"checkDigit,omitempty"`
	RestrictedTicketIndicator  int                                                  `json:"restrictedTicketIndicator,omitempty"`
	TransactionType            int                                                  `json:"transactionType,omitempty"`
	ExtendedPaymentCode        string                                               `json:"extendedPaymentCode,omitempty"`
	PassengerName              string                                               `json:"passengerName,omitempty"`
	CustomerCode               string                                               `json:"customerCode,omitempty"`
	DocumentType               string                                               `json:"documentType,omitempty"`
	DocumentNumber             string                                               `json:"documentNumber,omitempty"`
	DocumentNumberOfParts      int                                                  `json:"documentNumberOfParts,omitempty"`
	InvoiceNumber              string                                               `json:"invoiceNumber,omitempty"`
	InvoiceDate                int                                                  `json:"invoiceDate,omitempty"`
	AdditionalCharges          string                                               `json:"additionalCharges,omitempty"`
	TotalFeeAmount             string                                               `json:"totalFeeAmount,omitempty"`
	ClearingSequence           string                                               `json:"clearingSequence,omitempty"`
	ClearingCount              string                                               `json:"clearingCount,omitempty"`
	TotalClearingAmount        string                                               `json:"totalClearingAmount,omitempty"`
	NumberOfPassengers         int                                                  `json:"numberOfPassengers,omitempty"`
	ReservationSystemCode      string                                               `json:"reservationSystemCode,omitempty"`
	ProcessIdentifier          string                                               `json:"processIdentifier,omitempty"`
	TicketIssueDate            string                                               `json:"ticketIssueDate,omitempty"`
	ElectronicTicketIndicator  bool                                                 `json:"electronicTicketIndicator,omitempty"`
	OriginalTicketNumber       string                                               `json:"originalTicketNumber,omitempty"`
	PurchaseType               string                                               `json:"purchaseType,omitempty"`
	CreditReasonIndicator      string                                               `json:"creditReasonIndicator,omitempty"`
	TicketChangeIndicator      string                                               `json:"ticketChangeIndicator,omitempty"`
	PlanNumber                 string                                               `json:"planNumber,omitempty"`
	ArrivalDate                string                                               `json:"arrivalDate,omitempty"`
	RestrictedTicketDesciption string                                               `json:"restrictedTicketDesciption,omitempty"`
	ExchangeTicketAmount       string                                               `json:"exchangeTicketAmount,omitempty"`
	ExchangeTicketFeeAmount    string                                               `json:"exchangeTicketFeeAmount,omitempty"`
	ReservationType            string                                               `json:"reservationType,omitempty"`
	BoardingFeeAmount          string                                               `json:"boardingFeeAmount,omitempty"`
	Legs                       []*TravelInformationTransitAirlineLegs               `json:"legs,omitempty"`
	AncillaryInformation       *TravelInformationTransitAirlineAncillaryInformation `json:"ancillaryInformation,omitempty"`
}

type TravelInformationTransitAirlineTicketIssuer struct {
	Code               string `json:"code,omitempty"`
	Name               string `json:"name,omitempty"`
	Address            string `json:"address,omitempty"`
	Locality           string `json:"locality,omitempty"`
	AdministrativeArea string `json:"administrativeArea,omitempty"`
	PostalCode         string `json:"postalCode,omitempty"`
	Country            string `json:"country,omitempty"`
}

type TravelInformationTransitAirlineLegs struct {
	CarrierCode              string `json:"carrierCode,omitempty"`
	FlightNumber             string `json:"flightNumber,omitempty"`
	OriginatingAirportCode   string `json:"originatingAirportCode,omitempty"`
	Class                    string `json:"propertyClass,omitempty"`
	StopoverIndicator        int    `json:"stopoverIndicator,omitempty"`
	DepartureDate            int    `json:"departureDate,omitempty"`
	DestinationAirportCode   string `json:"destinationAirportCode,omitempty"`
	FareBasis                string `json:"fareBasis,omitempty"`
	DepartTaxAmount          string `json:"departTaxAmount,omitempty"`
	ConjunctionTicket        string `json:"conjunctionTicket,omitempty"`
	ExchangeTicketNumber     string `json:"exchangeTicketNumber,omitempty"`
	CouponNumber             string `json:"couponNumber,omitempty"`
	DepartureTime            int    `json:"departureTime,omitempty"`
	DepartureTimeMeridian    string `json:"departureTimeMeridian,omitempty"`
	ArrivalTime              int    `json:"arrivalTime,omitempty"`
	ArrivalTimeMeridian      string `json:"arrivalTimeMeridian,omitempty"`
	EndorsementsRestrictions string `json:"endorsementsRestrictions,omitempty"`
	TotalFareAmount          string `json:"totalFareAmount,omitempty"`
	FeeAmount                string `json:"feeAmount,omitempty"`
	TaxAmount                string `json:"taxAmount,omitempty"`
}

type TravelInformationTransitAirlineAncillaryInformation struct {
	TicketNumber          string                                                        `json:"ticketNumber,omitempty"`
	PassengerName         string                                                        `json:"passengerName,omitempty"`
	ConnectedTicketNumber string                                                        `json:"connectedTicketNumber,omitempty"`
	CreditReasonIndicator string                                                        `json:"creditReasonIndicator,omitempty"`
	Service               []*TravelInformationTransitAirlineAncillaryInformationService `json:"service,omitempty"`
}

type TravelInformationTransitAirlineAncillaryInformationService struct {
	CategoryCode    string `json:"categoryCode,omitempty"`
	SubCategoryCode string `json:"subCategoryCode,omitempty"`
}
