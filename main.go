package main

import (
	"fmt"
	"log"

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

	// Start REST API server to expose Ethereum parser functionality
	go restapi.StartServer(ethereumParser, memoryStorage)

	// Simulate fetching current block
	currentBlock, err := getCurrentBlock(rpcClient)
	if err != nil {
		log.Printf("Error fetching block number: %v\n", err)
		return
	}
	fmt.Printf("Current block: %d\n", currentBlock)

	select {} // Keep the main function running
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
