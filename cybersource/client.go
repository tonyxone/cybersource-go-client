package cybersource

type Client struct {
	merchantKeyId     string
	merchantSecretKey string
	merchantId        string
	requestHost       string
}

type Params struct {
	MerchantKeyId     string
	MerchantSecretKey string
	MerchantId        string
	RequestHost       string
}

func NewClient(params Params) *Client {
	return &Client{
		merchantKeyId:     params.MerchantKeyId,
		merchantSecretKey: params.MerchantSecretKey,
		merchantId:        params.MerchantId,
		requestHost:       params.RequestHost,
	}
}
