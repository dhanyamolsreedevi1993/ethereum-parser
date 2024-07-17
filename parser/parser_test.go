package parser

import (
	"testing"
)

func TestEthereumParser_GetCurrentBlock(t *testing.T) {
	parser := NewEthereumParser()
	expectedBlock := 0
	if block := parser.GetCurrentBlock(); block != expectedBlock {
		t.Errorf("GetCurrentBlock() returned %d, expected %d", block, expectedBlock)
	}
}

func TestEthereumParser_SubscribeAndGetTransactions(t *testing.T) {
	parser := NewEthereumParser()
	address := "0x1234567890abcdef"

	// Test Subscribe method
	subscribed := parser.Subscribe(address)
	if !subscribed {
		t.Errorf("Subscribe() returned false, expected true")
	}

	// Test GetTransactions method
	transactions := parser.GetTransactions(address)
	if len(transactions) != 1 {
		t.Errorf("GetTransactions() returned %d transactions, expected 1", len(transactions))
	}
	// Validate the dummy transaction details
	expectedTransaction := &Transaction{
		Hash:     "0xdummyhash",
		From:     "0xfromaddress",
		To:       address,
		Value:    "100",
		Gas:      "21000",
		GasPrice: "50",
	}
	if *transactions[0] != *expectedTransaction {
		t.Errorf("GetTransactions() returned %+v, expected %+v", transactions[0], expectedTransaction)
	}
}
