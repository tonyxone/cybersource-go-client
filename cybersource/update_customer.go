package cybersource

import (
	"encoding/json"
	"github.com/tonyxone/cybersource-rest-sdk-go/cybersource/model/customer/request"
	"github.com/tonyxone/cybersource-rest-sdk-go/cybersource/model/customer/response"
	"io/ioutil"
)

func (c *Client) UpdateCustomer(customerID string, req *request.CustomerRequest) (*response.CustomersResponse, error) {
	resource := "/tms/v2/customers/" + customerID
	resp, err := c.doPatch(resource, req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var customerResp response.CustomersResponse
	err = json.Unmarshal(body, &customerResp)
	if err != nil {
		return nil, err
	}
	return &customerResp, nil
}
