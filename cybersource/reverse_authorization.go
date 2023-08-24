package cybersource

import (
	"encoding/json"
	"github.com/tonyxone/cybersource-rest-sdk-go/cybersource/model/request"
	"github.com/tonyxone/cybersource-rest-sdk-go/cybersource/model/response"
	"io/ioutil"
)

func (c *Client) ReverseAuthorization(requestID string, req *request.ReverseAuthRequest) (*response.PaymentResponse, error) {
	resource := "/pts/v2/payments/" + requestID + "/reversals"
	resp, err := c.doPost(resource, req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var paymentResp response.PaymentResponse
	err = json.Unmarshal(body, &paymentResp)
	if err != nil {
		return nil, err
	}
	return &paymentResp, nil
}
