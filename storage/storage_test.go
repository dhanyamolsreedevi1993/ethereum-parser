package storage

import (
	"testing"

	"github.com/dhanyamolsreedevi1993/ethereum-parser/parser"
)

func TestMemoryStorage_SaveTransactionAndGetTransactions(t *testing.T) {
	// Initialize MemoryStorage
	storage := NewMemoryStorage()

	// Test SaveTransaction method
	address := "0x1234567890abcdef"
	tx := &parser.Transaction{
		Hash:     "0xtransactionhash",
		From:     "0xfromaddress",
		To:       address,
		Value:    "50",
		Gas:      "21000",
		GasPrice: "100",
	}
	storage.SaveTransaction(address, tx)

	// Test GetTransactions method
	transactions := storage.GetTransactions(address)
	if len(transactions) != 1 {
		t.Errorf("GetTransactions() returned %d transactions, expected 1", len(transactions))
	}
	if *transactions[0] != *tx {
		t.Errorf("GetTransactions() returned %+v, expected %+v", transactions[0], tx)
	}
}
