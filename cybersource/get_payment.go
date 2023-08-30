package cybersource

import (
	"encoding/json"
	"github.com/tonyxone/cybersource-rest-sdk-go/cybersource/model/payment/response"
	"io/ioutil"
)

func (c *Client) GetPayment(requestID string) (*response.GetPaymentResponse, error) {
	resource := "/pts/v2/payments/" + requestID
	resp, err := c.doGet(resource)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var getPaymentResp response.GetPaymentResponse
	err = json.Unmarshal(body, &getPaymentResp)
	if err != nil {
		return nil, err
	}
	return &getPaymentResp, nil
}
