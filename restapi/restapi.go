package restapi

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dhanyamolsreedevi1993/ethereum-parser/parser"
	"github.com/dhanyamolsreedevi1993/ethereum-parser/storage"
)

var (
	ethereumParser *parser.EthereumParser
	memoryStorage  *storage.MemoryStorage
)

func init() {
	ethereumParser = parser.NewEthereumParser()
	memoryStorage = storage.NewMemoryStorage()
}

func StartServer(parserInstance *parser.EthereumParser, storageInstance *storage.MemoryStorage) {
	ethereumParser = parserInstance
	memoryStorage = storageInstance

	http.HandleFunc("/subscribe", handleSubscribe)
	http.HandleFunc("/transactions", handleTransactions)

	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleSubscribe(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	address := r.FormValue("address")
	if address == "" {
		http.Error(w, "Address parameter is required", http.StatusBadRequest)
		return
	}

	success := ethereumParser.Subscribe(address)
	if !success {
		http.Error(w, "Failed to subscribe to address", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Subscribed successfully"))
}

func handleTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	address := r.FormValue("address")
	if address == "" {
		http.Error(w, "Address parameter is required", http.StatusBadRequest)
		return
	}

	transactions := memoryStorage.GetTransactions(address)
	jsonResponse, err := json.Marshal(transactions)
	if err != nil {
		http.Error(w, "Failed to marshal transactions", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
