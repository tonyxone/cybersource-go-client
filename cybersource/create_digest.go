package cybersource

import (
	"crypto/sha256"
	"encoding/base64"
)

func (c *Client) createDigest(payload []byte) string {
	hash := sha256.Sum256(payload)
	base64EncodedHash := base64.StdEncoding.EncodeToString(hash[:])
	digest := "SHA-256=" + base64EncodedHash
	return digest
}
