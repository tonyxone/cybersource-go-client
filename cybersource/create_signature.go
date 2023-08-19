package cybersource

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/http"
)

func (c *Client) createSignatureHeader(httpMethod string, requestTarget string, digest string, gmtDateTime string) string {
	var sigHeader string

	// KeyId is the key obtained from EBC
	sigHeader += "keyid=\"" + c.merchantKeyId + "\""

	// Algorithm should be always HmacSHA256 for http signature
	sigHeader += ", algorithm=\"HmacSHA256\""

	/* Headers - list is choosen based on HTTP method. Digest is not required for GET Method */
	getHeaders := "host date (request-target)" + " " + "v-c-merchant-id"
	postHeaders := "host date (request-target) digest v-c-merchant-id"

	if httpMethod == "GET" {
		sigHeader += ", headers=\"" + getHeaders + "\""
	} else if httpMethod == "POST" {
		sigHeader += ", headers=\"" + postHeaders + "\""
	}

	// Get Value for paramter 'Signature' to be passed to Signature Header
	signatureValue := c.createSignatureParam(httpMethod, requestTarget, digest, gmtDateTime)
	sigHeader += ", signature=\"" + signatureValue + "\""

	return sigHeader
}

func (c *Client) createSignatureParam(httpMethod string, requestTarget, digest string, gmtDateTime string) string {
	/* This method returns value for paramter Signature which is then passed to Signature header
	 * paramter 'Signature' is calucated based on below key values and then signed with SECRET KEY -
	 * host: Sandbox (apitest.cybersource.com) or Production (api.cybersource.com) hostname
	 * date: "HTTP-date" format as defined by RFC7231.
	 * (request-target): Should be in format of httpMethod: path
	 *                   Example: "post /pts/v2/payments"
	 * Digest: Only needed for POST calls.
	 *          digestString = BASE64( HMAC-SHA256 ( Payload ));
	 *          Digest: "SHA-256=" + digestString;
	 * v-c-merchant-id: set value to Cybersource Merchant ID
	 *                   This ID can be found on EBC portal
	 */

	var signatureString string
	signatureString += "\n"
	signatureString += "host: " + c.requestHost + "\n"
	signatureString += "date: " + gmtDateTime + "\n"
	signatureString += "(request-target): "

	if httpMethod == http.MethodGet {
		signatureString += "get " + requestTarget
	} else if httpMethod == http.MethodPost {
		signatureString += "post " + requestTarget
	}

	signatureString += "\n"

	if httpMethod == http.MethodPost {
		signatureString += "digest: " + digest + "\n"
	}

	signatureString += "v-c-merchant-id: " + c.merchantId
	if len(signatureString) > 1 {
		signatureString = signatureString[1:]
	}

	/* Signature string generated from above parameters is Signed with SecretKey hased with SHA256 and base64 encoded.
	 *  Secret Key is Base64 decoded before signing
	 */
	decodedKey, err := base64.StdEncoding.DecodeString(c.merchantSecretKey)
	if err != nil {
		panic(err)
	}

	hmacKey := hmac.New(sha256.New, decodedKey)
	_, err = hmacKey.Write([]byte(signatureString))
	if err != nil {
		panic(err)
	}

	signatureBytes := hmacKey.Sum(nil)
	base64EncodedSignature := base64.StdEncoding.EncodeToString(signatureBytes)
	return base64EncodedSignature
}
