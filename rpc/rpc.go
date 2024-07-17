package rpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	ethereumRPCURL = "https://cloudflare-eth.com"
)

type EthereumRPCClient struct {
	url string
}

func NewEthereumRPCClient() *EthereumRPCClient {
	return &EthereumRPCClient{
		url: ethereumRPCURL,
	}
}

func (c *EthereumRPCClient) Call(method string, params []interface{}, result interface{}) error {
	requestBody, err := json.Marshal(map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  method,
		"params":  params,
		"id":      1,
	})
	if err != nil {
		return err
	}

	resp, err := http.Post(c.url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return err
	}

	return nil
}
