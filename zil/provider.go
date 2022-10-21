package zil

import (
	"github.com/go-resty/resty/v2"
)

type Provider struct {
	host string
}

func NewProvider(host string) *Provider {
	return &Provider{host: host}
}

func (provider *Provider) GetSmartContractSubState(contractAddress string, params ...interface{}) (string, error) {
	type request struct {
		Id      string      `json:"id"`
		Jsonrpc string      `json:"jsonrpc"`
		Method  string      `json:"method"`
		Params  interface{} `json:"params"`
	}

	requestParams := append(
		[]interface{}{
			contractAddress,
		}, params...,
	)

	resp, err := resty.New().R().
		SetBody(&request{
			Id:      "1",
			Jsonrpc: "2.0",
			Method:  "GetSmartContractSubState",
			Params:  requestParams,
		}).
		Post(provider.host)
	if err != nil {
		return "", err
	}

	return string(resp.Body()), nil

}
