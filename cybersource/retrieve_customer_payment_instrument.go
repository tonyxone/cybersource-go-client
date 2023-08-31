package cybersource

import (
	"encoding/json"
	"github.com/tonyxone/cybersource-rest-sdk-go/cybersource/model/token_management/response"
	"io/ioutil"
)

func (c *Client) RetrieveCustomerPaymentInstrument(customerID string, paymentInstrumentTokenID string) (*response.CustomerPaymentInstrumentResponse, error) {
	resource := "/tms/v2/customers/" + customerID + "/payment-instruments/" + paymentInstrumentTokenID
	resp, err := c.doGet(resource)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var instrumentResp response.CustomerPaymentInstrumentResponse
	err = json.Unmarshal(body, &instrumentResp)
	if err != nil {
		return nil, err
	}
	return &instrumentResp, nil
}
