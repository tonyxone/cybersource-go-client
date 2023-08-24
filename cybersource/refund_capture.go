package cybersource

import (
	"encoding/json"
	"github.com/tonyxone/cybersource-rest-sdk-go/cybersource/model/request"
	"github.com/tonyxone/cybersource-rest-sdk-go/cybersource/model/response"
	"io/ioutil"
)

func (c *Client) RefundCapture(requestID string, req *request.RefundRequest) (*response.RefundResponse, error) {
	resource := "/pts/v2/captures/" + requestID + "/refunds"
	resp, err := c.doPost(resource, req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var refundResp response.RefundResponse
	err = json.Unmarshal(body, &refundResp)
	if err != nil {
		return nil, err
	}
	return &refundResp, nil
}
