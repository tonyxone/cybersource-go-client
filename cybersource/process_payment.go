package cybersource

import (
	"encoding/json"
	"fmt"
	"github.com/tonyxone/cybersource-rest-sdk-go/cybersource/model/request"
	"github.com/tonyxone/cybersource-rest-sdk-go/cybersource/model/response"
	"io/ioutil"
)

func (c *Client) ProcessPayment(createPaymentReq *request.CreatePaymentRequest) (*response.CreatePaymentResponse, error) {
	resource := "/pts/v2/payments"
	resp, err := c.doPost(resource, createPaymentReq)
	fmt.Println(resp.StatusCode)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var createPaymentResponse response.CreatePaymentResponse
	err = json.Unmarshal(body, &createPaymentResponse)
	if err != nil {
		return nil, err
	}
	return &createPaymentResponse, nil
}
