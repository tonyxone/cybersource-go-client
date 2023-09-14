package cybersource

import (
	"encoding/json"
	"github.com/tonyxone/cybersource-rest-sdk-go/cybersource/model/token_management/request"
	"github.com/tonyxone/cybersource-rest-sdk-go/cybersource/model/token_management/response"
	"io/ioutil"
	"net/http"
)

func (c *Client) GenerateCaptureContext(req *request.GenerateCaptureContextRequest) (*response.GenerateCaptureContextResponse, error) {
	resource := "/microform/v2/sessions"
	resp, err := c.doPost(resource, req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var captureContextResp response.GenerateCaptureContextResponse
	if resp.StatusCode != http.StatusCreated {
		err = json.Unmarshal(body, &captureContextResp)
		if err != nil {
			return nil, err
		}
	} else {
		captureContextResp.CaptureContext = string(body)
	}

	return &captureContextResp, nil
}
