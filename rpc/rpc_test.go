package rpc

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEthereumRPCClient_Call(t *testing.T) {
	// Mock server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Validate the request
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST method, got %s", r.Method)
		}
		if r.URL.String() != ethereumRPCURL {
			t.Errorf("Expected URL %s, got %s", ethereumRPCURL, r.URL.String())
		}

		// Respond with dummy JSON RPC result
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"result":"0x4b7"}`))
	}))

	defer mockServer.Close()

	// Test EthereumRPCClient Call method
	client := NewEthereumRPCClient()
	client.url = mockServer.URL

	var result struct {
		Result string `json:"result"`
	}
	err := client.Call("eth_blockNumber", []interface{}{}, &result)
	if err != nil {
		t.Errorf("Call() returned error: %v", err)
	}
	if result.Result != "0x4b7" {
		t.Errorf("Call() returned result %s, expected 0x4b7", result.Result)
	}
}
