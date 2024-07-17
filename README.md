# Ethereum Blockchain Parser

## Introduction

The Ethereum Blockchain Parser is a system designed to monitor and query transactions for specific Ethereum addresses using JSON-RPC calls to an Ethereum node. It includes components for handling blockchain interaction, data storage, and a REST API for external access.

### Project Goals

- **Automated Testing**: Ensure robustness and correctness through automated tests.
- **Error Handling**: Handle errors gracefully across all components.
- **Good Practices**: Follow Go language best practices for structuring, modularizing, and documenting code.
- **Project Structure**: Maintain a clear separation of concerns and modular components.
- **Modularization**: Implement clear interfaces between components to support future expansions.
- **Logging**: Provide informative logging throughout the application.
- **Documentation**: Comprehensive documentation for easy understanding and usage.

## Components

### 1. **parser Package**

The `parser` package defines structures and methods related to interacting with Ethereum transactions and addresses.

- **Transaction**: Represents an Ethereum transaction with fields for hash, sender, receiver, value, gas, and gas price.
  
- **Parser Interface**: Specifies methods for retrieving the current block, subscribing to an address, and fetching transactions for an address.
  
- **EthereumParser Struct**: Implements the `Parser` interface. Manages subscription states and transaction storage with thread-safe access using a mutex.

### 2. **rpc Package**

The `rpc` package provides an interface to interact with an Ethereum node using JSON-RPC.

- **EthereumRPCClient Struct**: Handles HTTP requests to the Ethereum node's JSON-RPC endpoint. Supports making arbitrary JSON-RPC calls and decoding responses.

### 3. **storage Package**

The `storage` package manages the persistence of transaction data.

- **Storage Interface**: Defines methods for getting transactions and saving transactions for an address.
  
- **MemoryStorage Struct**: Implements the `Storage` interface using an in-memory map. Suitable for testing and small-scale applications.

### 4. **restapi Package**

The `restapi` package exposes the Ethereum parser functionalities via HTTP endpoints.

- **StartServer Function**: Initializes an HTTP server exposing `/subscribe` and `/transactions` endpoints.
  
- **handleSubscribe Function**: Handles POST requests to `/subscribe` endpoint for subscribing to an Ethereum address.
  
- **handleTransactions Function**: Handles GET requests to `/transactions` endpoint for fetching transactions of a subscribed address.

### 5. **main.go**

The `main` package initializes and starts the Ethereum parser system.

- **main Function**: Initializes instances of `EthereumParser`, `EthereumRPCClient`, and `MemoryStorage`. Starts the REST API server and fetches the current Ethereum block number periodically.

- **getCurrentBlock Function**: Fetches the current Ethereum block number using the `EthereumRPCClient`.

## Running the Application

### Prerequisites

- Go programming language installed (version 1.16+ recommended).
- Access to an Ethereum node's JSON-RPC endpoint (default: https://cloudflare-eth.com).

### Steps

1. **Clone the Repository**

   ```bash
   git clone <repository_url>
   cd ethereum-parser

2. **Build and Run**

   ```bash
   go build -o ethereum-parser main.go
   ./ethereum-parser

   This will start the application, initializing the Ethereum parser and exposing the REST API server on port 8080.

3. **Interact via REST API**

   - Subscribe to Address: Send a POST request to http://localhost:8080/subscribe with address parameter.

    Example
    ```bash
        curl -X POST 'http://localhost:8080/subscribe' --data 'address=0x1234567890abcdef'

        or 

        Send a POST request in postman : http://localhost:8080/subscribe?address='address=0x1234567890abcdef' 

 - Fetch Transactions: Send a GET request to http://localhost:8080/transactions with address parameter.
    Example
    ```bash
        curl -X GET 'http://localhost:8080/transactions?address=0x1234567890abcdef'
        or 
        Send a GET request in postman : http://localhost:8080/transactions?address='address=0x1234567890abcdef' 
    
    - If the adress is not subscribed ,transaction request will fail,subscribe first and send transaction request.

## Testing

### Running Tests

  **To run automated tests for the project:**
    ```bash
        go test ./...
        -This command runs tests for all packages (parser, rpc, storage, restapi) and ensures functionality is correct across components.

## Conclusion

    - The Ethereum Blockchain Parser provides a scalable solution for monitoring and querying Ethereum transactions. It leverages     modular design principles, adheres to best practices, and offers extensive testing to ensure reliability and maintainability.

