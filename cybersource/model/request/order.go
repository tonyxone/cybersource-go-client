package request

type OrderInformation struct {
	AmountDetails            *OrderInformationAmountDetails   `json:"amountDetails,omitempty"`
	BillTo                   *OrderInformationBillTo          `json:"billTo,omitempty"`
	ShipTo                   *OrderInformationShipTo          `json:"shipTo,omitempty"`
	LineItems                []*OrderInformationLineItems     `json:"lineItems,omitempty"`
	InvoiceDetails           *OrderInformationInvoiceDetails  `json:"invoiceDetails,omitempty"`
	ShippingDetails          *OrderInformationShippingDetails `json:"shippingDetails,omitempty"`
	ReturnsAccepted          bool                             `json:"returnsAccepted,omitempty"`
	IsCryptocurrencyPurchase string                           `json:"isCryptocurrencyPurchase,omitempty"`
	PreOrder                 string                           `json:"preOrder,omitempty"`
	PreOrderDate             string                           `json:"preOrderDate,omitempty"`
	Reordered                bool                             `json:"reordered,omitempty"`
	TotalOffersCount         string                           `json:"totalOffersCount,omitempty"`
}

type OrderInformationAmountDetails struct {
	TotalAmount             string                                                `json:"totalAmount,omitempty"`
	SubTotalAmount          string                                                `json:"subTotalAmount,omitempty"`
	Currency                string                                                `json:"currency,omitempty"`
	DiscountAmount          string                                                `json:"discountAmount,omitempty"`
	DutyAmount              string                                                `json:"dutyAmount,omitempty"`
	GratuityAmount          string                                                `json:"gratuityAmount,omitempty"`
	TaxAmount               string                                                `json:"taxAmount,omitempty"`
	NationalTaxIncluded     string                                                `json:"nationalTaxIncluded,omitempty"`
	TaxAppliedAfterDiscount string                                                `json:"taxAppliedAfterDiscount,omitempty"`
	TaxAppliedLevel         string                                                `json:"taxAppliedLevel,omitempty"`
	TaxTypeCode             string                                                `json:"taxTypeCode,omitempty"`
	FreightAmount           string                                                `json:"freightAmount,omitempty"`
	ForeignAmount           string                                                `json:"foreignAmount,omitempty"`
	ForeignCurrency         string                                                `json:"foreignCurrency,omitempty"`
	ExchangeRate            string                                                `json:"exchangeRate,omitempty"`
	ExchangeRateTimeStamp   string                                                `json:"exchangeRateTimeStamp,omitempty"`
	Surcharge               *OrderInformationAmountDetailsSurcharge               `json:"surcharge,omitempty"`
	SettlementAmount        string                                                `json:"settlementAmount,omitempty"`
	SettlementCurrency      string                                                `json:"settlementCurrency,omitempty"`
	AmexAdditionalAmounts   []*OrderInformationAmountDetailsAmexAdditionalAmounts `json:"amexAdditionalAmounts,omitempty"`
	TaxDetails              []*OrderInformationAmountDetailsTaxDetails            `json:"taxDetails,omitempty"`
	ServiceFeeAmount        string                                                `json:"serviceFeeAmount,omitempty"`
	OriginalAmount          string                                                `json:"originalAmount,omitempty"`
	OriginalCurrency        string                                                `json:"originalCurrency,omitempty"`
	CashbackAmount          string                                                `json:"cashbackAmount,omitempty"`
	CurrencyConversion      *OrderInformationAmountDetailsCurrencyConversion      `json:"currencyConversion,omitempty"`
}

type OrderInformationAmountDetailsSurcharge struct {
	Amount      string `json:"amount,omitempty"`
	Description string `json:"description,omitempty"`
}

type OrderInformationAmountDetailsAmexAdditionalAmounts struct {
	Code   string `json:"code,omitempty"`
	Amount string `json:"amount,omitempty"`
}

type OrderInformationAmountDetailsTaxDetails struct {
	Type          string `json:"type,omitempty"`
	Amount        string `json:"amount,omitempty"`
	Rate          string `json:"rate,omitempty"`
	Code          string `json:"code,omitempty"`
	TaxID         string `json:"taxId,omitempty"`
	Applied       bool   `json:"applied,omitempty"`
	ExemptionCode string `json:"exemptionCode,omitempty"`
}

type OrderInformationAmountDetailsCurrencyConversion struct {
	Indicator        string `json:"indicator,omitempty"`
	ReconciliationId string `json:"reconciliationId,omitempty"`
	ID               string `json:"id,omitempty"`
}

type OrderInformationBillTo struct {
	FirstName            string                         `json:"firstName,omitempty"`
	LastName             string                         `json:"lastName,omitempty"`
	MiddleName           string                         `json:"middleName,omitempty"`
	NameSuffix           string                         `json:"nameSuffix,omitempty"`
	Title                string                         `json:"title,omitempty"`
	Company              *OrderInformationBillToCompany `json:"company,omitempty"`
	Address1             string                         `json:"address1,omitempty"`
	Address2             string                         `json:"address2,omitempty"`
	Address3             string                         `json:"address3,omitempty"`
	Address4             string                         `json:"address4,omitempty"`
	Locality             string                         `json:"locality,omitempty"`
	AdministrativeArea   string                         `json:"administrativeArea,omitempty"`
	PostalCode           string                         `json:"postalCode,omitempty"`
	County               string                         `json:"county,omitempty"`
	Country              string                         `json:"country,omitempty"`
	District             string                         `json:"district,omitempty"`
	BuildingNumber       string                         `json:"buildingNumber,omitempty"`
	Email                string                         `json:"email,omitempty"`
	EmailDomain          string                         `json:"emailDomain,omitempty"`
	PhoneNumber          string                         `json:"phoneNumber,omitempty"`
	PhoneType            string                         `json:"phoneType,omitempty"`
	VerificationStatus   string                         `json:"verificationStatus,omitempty"`
	AlternatePhoneNumber string                         `json:"alternatePhoneNumber,omitempty"`
	AlternateEmail       string                         `json:"alternateEmail,omitempty"`
}

type OrderInformationBillToCompany struct {
	Name               string `json:"name,omitempty"`
	Address1           string `json:"address1,omitempty"`
	Address2           string `json:"address2,omitempty"`
	Locality           string `json:"locality,omitempty"`
	AdministrativeArea string `json:"administrativeArea,omitempty"`
	PostalCode         string `json:"postalCode,omitempty"`
	Country            string `json:"country,omitempty"`
}

type OrderInformationShipTo struct {
	Title              string `json:"title,omitempty"`
	FirstName          string `json:"firstName,omitempty"`
	MiddleName         string `json:"middleName,omitempty"`
	LastName           string `json:"lastName,omitempty"`
	Address1           string `json:"address1,omitempty"`
	Address2           string `json:"address2,omitempty"`
	Locality           string `json:"locality,omitempty"`
	AdministrativeArea string `json:"administrativeArea,omitempty"`
	PostalCode         string `json:"postalCode,omitempty"`
	Country            string `json:"country,omitempty"`
	District           string `json:"district,omitempty"`
	BuildingNumber     string `json:"buildingNumber,omitempty"`
	PhoneNumber        string `json:"phoneNumber,omitempty"`
	Company            string `json:"company,omitempty"`
	DestinationTypes   string `json:"destinationTypes,omitempty"`
	DestinationCode    int    `json:"destinationCode,omitempty"`
	Method             string `json:"method,omitempty"`
}

type OrderInformationLineItems struct {
	ProductCode               string                                     `json:"productCode,omitempty"`
	ProductName               string                                     `json:"productName,omitempty"`
	ProductSku                string                                     `json:"productSku,omitempty"`
	Quantity                  int                                        `json:"quantity,omitempty"`
	UnitPrice                 string                                     `json:"unitPrice,omitempty"`
	UnitOfMeasure             string                                     `json:"unitOfMeasure,omitempty"`
	TotalAmount               string                                     `json:"totalAmount,omitempty"`
	TaxAmount                 string                                     `json:"taxAmount,omitempty"`
	TaxRate                   string                                     `json:"taxRate,omitempty"`
	TaxAppliedAfterDiscount   string                                     `json:"taxAppliedAfterDiscount,omitempty"`
	TaxStatusIndicator        string                                     `json:"taxStatusIndicator,omitempty"`
	TaxTypeCode               string                                     `json:"taxTypeCode,omitempty"`
	AmountIncludesTax         bool                                       `json:"amountIncludesTax,omitempty"`
	TypeOfSupply              string                                     `json:"typeOfSupply,omitempty"`
	CommodityCode             string                                     `json:"commodityCode,omitempty"`
	DiscountAmount            string                                     `json:"discountAmount,omitempty"`
	DiscountApplied           bool                                       `json:"discountApplied,omitempty"`
	DiscountRate              string                                     `json:"discountRate,omitempty"`
	InvoiceNumber             string                                     `json:"invoiceNumber,omitempty"`
	TaxDetails                []*OrderInformationAmountDetailsTaxDetails `json:"taxDetails,omitempty"`
	FulfillmentType           string                                     `json:"fulfillmentType,omitempty"`
	Weight                    string                                     `json:"weight,omitempty"`
	WeightIdentifier          string                                     `json:"weightIdentifier,omitempty"`
	WeightUnit                string                                     `json:"weightUnit,omitempty"`
	ReferenceDataCode         string                                     `json:"referenceDataCode,omitempty"`
	ReferenceDataNumber       string                                     `json:"referenceDataNumber,omitempty"`
	ProductDescription        string                                     `json:"productDescription,omitempty"`
	GiftCardCurrency          int                                        `json:"giftCardCurrency,omitempty"`
	ShippingDestinationTypes  string                                     `json:"shippingDestinationTypes,omitempty"`
	Gift                      bool                                       `json:"gift,omitempty"`
	Passenger                 *OrderInformationPassenger                 `json:"passenger,omitempty"`
	AllowedExportCountries    []string                                   `json:"allowedExportCountries,omitempty"`
	RestrictedExportCountries []string                                   `json:"restrictedExportCountries,omitempty"`
}

type OrderInformationPassenger struct {
	Type        string `json:"type,omitempty"`
	Status      string `json:"status,omitempty"`
	Phone       string `json:"phone,omitempty"`
	FirstName   string `json:"firstName,omitempty"`
	LastName    string `json:"lastName,omitempty"`
	ID          string `json:"id,omitempty"`
	Email       string `json:"email,omitempty"`
	Nationality string `json:"nationality,omitempty"`
}

type OrderInformationInvoiceDetails struct {
	InvoiceNumber             string                                                     `json:"invoiceNumber,omitempty"`
	BarcodeNumber             string                                                     `json:"barcodeNumber,omitempty"`
	ExpirationDate            string                                                     `json:"expirationDate,omitempty"`
	PurchaseOrderNumber       string                                                     `json:"purchaseOrderNumber,omitempty"`
	PurchaseOrderDate         string                                                     `json:"purchaseOrderDate,omitempty"`
	PurchaseContactName       string                                                     `json:"purchaseContactName,omitempty"`
	Taxable                   bool                                                       `json:"taxable,omitempty"`
	VatInvoiceReferenceNumber string                                                     `json:"vatInvoiceReferenceNumber,omitempty"`
	CommodityCode             string                                                     `json:"commodityCode,omitempty"`
	MerchandiseCode           int                                                        `json:"merchandiseCode,omitempty"`
	TransactionAdviceAddendum []*OrderInformationInvoiceDetailsTransactionAdviceAddendum `json:"transactionAdviceAddendum,omitempty"`
	ReferenceDataCode         string                                                     `json:"referenceDataCode,omitempty"`
	ReferenceDataNumber       string                                                     `json:"referenceDataNumber,omitempty"`
	SalesSlipNumber           int                                                        `json:"salesSlipNumber,omitempty"`
	InvoiceDate               string                                                     `json:"invoiceDate,omitempty"`
	CostCenter                string                                                     `json:"costCenter,omitempty"`
	IssuerMessage             string
}

type OrderInformationInvoiceDetailsTransactionAdviceAddendum struct {
	Data string `json:"data,omitempty"`
}

type OrderInformationShippingDetails struct {
	GiftWrap           bool   `json:"giftWrap,omitempty"`
	ShippingMethod     string `json:"shippingMethod,omitempty"`
	ShipFromPostalCode string `json:"shipFromPostalCode,omitempty"`
}
