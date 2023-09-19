package cybersource

import (
	"encoding/json"
	"github.com/tonyxone/cybersource-rest-sdk-go/cybersource/model/payment/response"
	"io/ioutil"
)

func (c *Client) RetrieveTransaction(requestID string) (*response.TransactionDetails, error) {
	resource := "/tss/v2/transactions/" + requestID
	resp, err := c.doGet(resource)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var transactionDetails response.TransactionDetails
	err = json.Unmarshal(body, &transactionDetails)
	if err != nil {
		return nil, err
	}
	return &transactionDetails, nil
}
