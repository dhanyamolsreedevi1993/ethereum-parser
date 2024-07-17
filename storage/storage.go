package storage

import (
	"sync"

	"github.com/dhanyamolsreedevi1993/ethereum-parser/parser"
)

type MemoryStorage struct {
	mu           sync.Mutex
	transactions map[string][]*parser.Transaction
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		transactions: make(map[string][]*parser.Transaction),
	}
}

func (s *MemoryStorage) GetCurrentBlock() int {
	return 0 // Not implemented in memory storage
}

func (s *MemoryStorage) GetTransactions(address string) []*parser.Transaction {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.transactions[address]; !ok {
		return []*parser.Transaction{}
	}

	// Return a copy of the transactions to avoid modifying the original data
	transactions := make([]*parser.Transaction, len(s.transactions[address]))
	copy(transactions, s.transactions[address])
	return transactions
}

func (s *MemoryStorage) SaveTransaction(address string, tx *parser.Transaction) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.transactions[address] = append(s.transactions[address], tx)
}
