package cybersource

import (
	"encoding/json"
	"github.com/tonyxone/cybersource-rest-sdk-go/cybersource/model/payment/request"
	"github.com/tonyxone/cybersource-rest-sdk-go/cybersource/model/payment/response"
	"io/ioutil"
)

func (c *Client) ProcessCredit(req *request.CreateCreditRequest) (*response.CreditResponse, error) {
	resource := "/pts/v2/credits"
	resp, err := c.doPost(resource, req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var creditResp response.CreditResponse
	err = json.Unmarshal(body, &creditResp)
	if err != nil {
		return nil, err
	}
	return &creditResp, nil
}
