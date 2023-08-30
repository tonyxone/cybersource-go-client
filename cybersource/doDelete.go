package cybersource

import (
	"net/http"
	"time"
)

func (c *Client) doDelete(resource string) (*http.Response, error) {
	gmtDateTime := time.Now().UTC().String()
	url := "https://" + c.requestHost + resource

	digest := "" // Digest is not required for GET calls
	signature := c.createSignatureHeader(http.MethodDelete, resource, digest, gmtDateTime)

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return &http.Response{}, err
	}

	req.Header.Set("v-c-merchant-id", c.merchantId)
	req.Header.Set("Date", gmtDateTime)
	req.Header.Set("Host", c.requestHost)
	req.Header.Set("Signature", signature)

	httpClient := http.Client{}
	return httpClient.Do(req)
}
