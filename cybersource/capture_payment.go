package cybersource

import (
	"encoding/json"
	"github.com/tonyxone/cybersource-rest-sdk-go/cybersource/model/payment/request"
	"github.com/tonyxone/cybersource-rest-sdk-go/cybersource/model/payment/response"
	"io/ioutil"
)

func (c *Client) CapturePayment(requestID string, req *request.CapturePaymentRequest) (*response.CaptureResponse, error) {
	resource := "/pts/v2/payments/" + requestID + "/captures"
	resp, err := c.doPost(resource, req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var capturePaymentResp response.CaptureResponse
	err = json.Unmarshal(body, &capturePaymentResp)
	if err != nil {
		return nil, err
	}
	return &capturePaymentResp, nil
}
