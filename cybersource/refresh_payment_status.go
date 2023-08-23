package cybersource

import (
	"encoding/json"
	"github.com/tonyxone/cybersource-rest-sdk-go/cybersource/model/request"
	"github.com/tonyxone/cybersource-rest-sdk-go/cybersource/model/response"
	"io/ioutil"
)

func (c *Client) RefreshPaymentStatus(requestID string, request *request.RefreshPaymentStatusRequest) (*response.RefreshPaymentStatusResponse, error) {
	resource := "/pts/v2/refresh-payment-status/" + requestID
	resp, err := c.doPost(resource, request)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var refreshPaymentStatusResp response.RefreshPaymentStatusResponse
	err = json.Unmarshal(body, &refreshPaymentStatusResp)
	if err != nil {
		return nil, err
	}
	return &refreshPaymentStatusResp, nil
}
