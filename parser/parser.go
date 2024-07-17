package parser

import (
	"sync"
)

type Transaction struct {
	Hash     string
	From     string
	To       string
	Value    string
	Gas      string
	GasPrice string
}

type Parser interface {
	GetCurrentBlock() int
	Subscribe(address string) bool
	GetTransactions(address string) []*Transaction
}

type EthereumParser struct {
	currentBlock int
	subscribed   map[string]bool
	transactions map[string][]*Transaction
	mu           sync.Mutex
}

func NewEthereumParser() *EthereumParser {
	return &EthereumParser{
		subscribed:   make(map[string]bool),
		transactions: make(map[string][]*Transaction),
	}
}

func (ep *EthereumParser) GetCurrentBlock() int {
	return ep.currentBlock
}

func (ep *EthereumParser) Subscribe(address string) bool {
	ep.mu.Lock()
	defer ep.mu.Unlock()

	if _, ok := ep.subscribed[address]; ok {
		return false // Already subscribed
	}

	ep.subscribed[address] = true

	// Add a dummy transaction for testing purposes
	dummyTx := &Transaction{
		Hash:     "0xdummyhash",
		From:     "0xfromaddress",
		To:       address,
		Value:    "100",
		Gas:      "21000",
		GasPrice: "50",
	}
	ep.transactions[address] = append(ep.transactions[address], dummyTx)

	return true
}

func (ep *EthereumParser) GetTransactions(address string) []*Transaction {
	ep.mu.Lock()
	defer ep.mu.Unlock()

	if !ep.subscribed[address] {
		return nil // Address not subscribed
	}

	// Return transactions only if the address is subscribed
	return ep.transactions[address]
}
