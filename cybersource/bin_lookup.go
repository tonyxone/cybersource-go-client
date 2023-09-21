package cybersource

import (
	"encoding/json"
	"github.com/tonyxone/cybersource-rest-sdk-go/cybersource/model/payment/request"
	"github.com/tonyxone/cybersource-rest-sdk-go/cybersource/model/payment/response"
	"io/ioutil"
)

func (c *Client) LookupBIN(req *request.BINLookupRequest) (*response.BINLookupResponse, error) {
	resource := "/bin/v1/binlookup"
	resp, err := c.doPost(resource, req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var BINResp response.BINLookupResponse
	err = json.Unmarshal(body, &BINResp)
	if err != nil {
		return nil, err
	}
	return &BINResp, nil
}
