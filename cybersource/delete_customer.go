package cybersource

import (
	"errors"
	"io/ioutil"
)

func (c *Client) DeleteCustomer(customerID string) error {
	resource := "/tms/v2/customers/" + customerID
	resp, err := c.doDelete(resource)
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	if resp.StatusCode != 204 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(string(body))
	}

	return nil
}
