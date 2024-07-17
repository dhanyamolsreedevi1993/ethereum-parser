package main

import (
	"fmt"

	"github.com/dhanyamolsreedevi1993/ethereum-parser/parser"
	"github.com/dhanyamolsreedevi1993/ethereum-parser/restapi"
	"github.com/dhanyamolsreedevi1993/ethereum-parser/rpc"
	"github.com/dhanyamolsreedevi1993/ethereum-parser/storage"
)

func main() {
	// Initialize Ethereum parser, RPC client, and storage
	ethereumParser := parser.NewEthereumParser()
	rpcClient := rpc.NewEthereumRPCClient()
	memoryStorage := storage.NewMemoryStorage()

	// Simulate subscribing to an address
	address := "0x1234567890abcdef"
	fmt.Printf("Subscribing to address %s\n", address)
	ethereumParser.Subscribe(address)

	// Simulate fetching current block
	currentBlock, err := getCurrentBlock(rpcClient)
	if err != nil {
		fmt.Printf("Error fetching block number: %v\n", err)
		return
	}
	fmt.Printf("Current block: %d\n", currentBlock)

	// Simulate fetching transactions for the subscribed address
	transactions := ethereumParser.GetTransactions(address)
	fmt.Printf("Transactions for address %s:\n", address)
	for _, tx := range transactions {
		fmt.Printf("Hash: %s, From: %s, To: %s, Value: %s, Gas: %s, GasPrice: %s\n",
			tx.Hash, tx.From, tx.To, tx.Value, tx.Gas, tx.GasPrice)
	}

	// Start REST API server to expose Ethereum parser functionality
	restapi.StartServer(ethereumParser, memoryStorage)
}

func getCurrentBlock(rpcClient *rpc.EthereumRPCClient) (int, error) {
	var blockNumberResult struct {
		Result string `json:"result"`
	}

	err := rpcClient.Call("eth_blockNumber", []interface{}{}, &blockNumberResult)
	if err != nil {
		return 0, fmt.Errorf("error fetching block number: %v", err)
	}

	// Convert hex block number to integer
	var blockNumber int
	_, err = fmt.Sscanf(blockNumberResult.Result, "0x%x", &blockNumber)
	if err != nil {
		return 0, fmt.Errorf("error converting block number: %v", err)
	}

	return blockNumber, nil
}
