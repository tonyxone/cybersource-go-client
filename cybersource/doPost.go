package cybersource

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

func (c *Client) doPost(resource string, payload interface{}) (*http.Response, error) {
	gmtDateTime := time.Now().UTC().String()
	url := "https://" + c.requestHost + resource

	payLoadBytes, err := json.Marshal(payload)
	if err != nil {
		return &http.Response{}, err
	}

	digest := c.createDigest(payLoadBytes)
	signature := c.createSignatureHeader(http.MethodPost, resource, digest, gmtDateTime)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payLoadBytes))
	if err != nil {
		return &http.Response{}, err
	}

	req.Header.Set("v-c-merchant-id", c.merchantId)
	req.Header.Set("Date", gmtDateTime)
	req.Header.Set("Host", c.requestHost)
	req.Header.Set("Digest", digest)
	req.Header.Set("Signature", signature)
	req.Header.Set("Content-Type", "application/json")

	httpClient := http.Client{}
	return httpClient.Do(req)
}
