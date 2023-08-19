package cybersource

import (
	"fmt"
	"github.com/tonyxone/cybersource-rest-sdk-go/cybersource/model/request"
	"io/ioutil"
)

func (c *Client) ProcessPayment(createPaymentReq *request.CreatePaymentRequest) error {
	resource := "/pts/v2/payments"
	resp, err := c.doPost(resource, createPaymentReq)
	fmt.Println(resp.StatusCode)
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	return nil
}
