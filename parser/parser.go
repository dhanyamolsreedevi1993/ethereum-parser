package parser

import (
	"github.com/dhanyamolsreedevi1993/ethereum-parser/storage_interface"
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
	IsSubscribed(address string) bool
	GetTransactions(address string, storage storage_interface.Storage) []*Transaction
}

type EthereumParser struct {
	currentBlock int
	subscribed   map[string]bool
}

func NewEthereumParser() *EthereumParser {
	return &EthereumParser{
		subscribed: make(map[string]bool),
	}
}

func (ep *EthereumParser) GetCurrentBlock() int {
	return ep.currentBlock // Not implemented in parser (consider fetching from RPC if needed)
}

func (ep *EthereumParser) Subscribe(address string) bool {
	ep.subscribed[address] = true
	return true
}

func (ep *EthereumParser) IsSubscribed(address string) bool {
	_, ok := ep.subscribed[address]
	return ok
}

func (ep *EthereumParser) GetTransactions(address string, storage storage_interface.Storage) []*Transaction {
	return storage.GetTransactions(address)
}
