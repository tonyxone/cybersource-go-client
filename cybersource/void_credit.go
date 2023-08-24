package cybersource

import (
	"encoding/json"
	"github.com/tonyxone/cybersource-rest-sdk-go/cybersource/model/request"
	"github.com/tonyxone/cybersource-rest-sdk-go/cybersource/model/response"
	"io/ioutil"
)

func (c *Client) VoidCredit(requestID string, req *request.VoidRequest) (*response.VoidResponse, error) {
	resource := "/pts/v2/credits/" + requestID + "/voids"
	resp, err := c.doPost(resource, req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var voidResp response.VoidResponse
	err = json.Unmarshal(body, &voidResp)
	if err != nil {
		return nil, err
	}
	return &voidResp, nil
}
