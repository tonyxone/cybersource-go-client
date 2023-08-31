package cybersource

import (
	"encoding/json"
	"github.com/tonyxone/cybersource-rest-sdk-go/cybersource/model/token_management/request"
	"github.com/tonyxone/cybersource-rest-sdk-go/cybersource/model/token_management/response"
	"io/ioutil"
)

func (c *Client) CreateCustomerPaymentInstrument(customerID string, req *request.CustomerPaymentInstrumentRequest) (*response.CustomerPaymentInstrumentResponse, error) {
	resource := "/tms/v2/customers/" + customerID + "/payment-instruments"
	resp, err := c.doPost(resource, req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var paymentInstrumentResp response.CustomerPaymentInstrumentResponse
	err = json.Unmarshal(body, &paymentInstrumentResp)
	if err != nil {
		return nil, err
	}
	return &paymentInstrumentResp, nil
}
