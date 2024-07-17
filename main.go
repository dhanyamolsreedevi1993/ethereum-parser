package main

import (
	"fmt"

	"github.com/dhanyamolsreedevi1993/ethereum-parser/parser"
	"github.com/dhanyamolsreedevi1993/ethereum-parser/restapi"
	"github.com/dhanyamolsreedevi1993/ethereum-parser/rpc"
	"github.com/dhanyamolsreedevi1993/ethereum-parser/storage"
)

func getCurrentBlock(client *rpc.EthereumRPCClient) (int, error) {
	// Define the Ethereum RPC method for getting the current block number
	const ethBlockNumberMethod = "eth_blockNumber"

	// Create request parameters
	params := []interface{}{}

	// Call the RPC method
	var blockNumber string
	err := client.Call(ethBlockNumberMethod, params, &blockNumber)
	if err != nil {
		return 0, fmt.Errorf("failed to get current block number: %v", err)
	}

	// Convert block number string to integer
	return parseIntHex(blockNumber)
}

func parseIntHex(hex string) (int, error) {
	// Implement logic to convert hexadecimal string to integer
	// You can use libraries like "github.com/ethereum/go-ethereum/common/hexutil"
	return 0, fmt.Errorf("parseIntHex not implemented")
}

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
	transactions := ethereumParser.GetTransactions(address, memoryStorage)
	fmt.Printf("Transactions for address %s:\n", address)
	for _, tx := range transactions {
		fmt.Printf("Hash: %s, From: %s, To: %s, Value: %s, Gas: %s, GasPrice: %s\n",
			tx.Hash, tx.From, tx.To, tx.Value, tx.Gas, tx.GasPrice)
	}

	// Start REST API server to expose Ethereum parser functionality
	restapi.StartServer(ethereumParser, memoryStorage)
}
