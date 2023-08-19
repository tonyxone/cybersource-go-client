package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/tonyxone/cybersource-rest-sdk-go/cybersource"
	"github.com/tonyxone/cybersource-rest-sdk-go/cybersource/model"
	"github.com/tonyxone/cybersource-rest-sdk-go/cybersource/model/request"
)

func main() {
	testProcessPayment()
}

func testProcessPayment() {
	client := cybersource.NewClient(cybersource.Params{
		MerchantKeyId:     "your_merchant_key_id",
		MerchantSecretKey: "your_merchant_secret_key",
		MerchantId:        "your_merchant_id",
		RequestHost:       "apitest.cybersource.com",
	})

	createPaymentReq := &request.CreatePaymentRequest{
		ClientReferenceInformation: &model.ClientReferenceInformation{
			Code: uuid.New().String(),
		},
		ProcessingInformation: &model.ProcessingInformation{
			CommerceIndicator: "internet",
		},
		OrderInformation: &model.OrderInformation{
			AmountDetails: &model.OrderInformationAmountDetails{
				TotalAmount: "60.00",
				Currency:    "USD",
			},
			BillTo: &model.OrderInformationBillTo{
				FirstName:          "John",
				LastName:           "Doe",
				Address1:           "201 S. Division St.",
				PostalCode:         "48104-2201",
				Locality:           "Ann Arbor",
				AdministrativeArea: "MI",
				Country:            "US",
				PhoneNumber:        "999999999",
				Email:              "test@cybs.com",
			},
		},
		PaymentInformation: &model.PaymentInformation{
			Card: &model.PaymentInformationCard{
				ExpirationYear:  "2031",
				Number:          "5555555555554444",
				SecurityCode:    "123",
				ExpirationMonth: "12",
				Type:            "002",
			},
		},
	}
	err := client.ProcessPayment(createPaymentReq)
	fmt.Println(err)
}
