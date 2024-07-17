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
	return 0
}

func (s *MemoryStorage) Subscribe(address string) bool {
	return true
}

func (s *MemoryStorage) GetTransactions(address string) []*parser.Transaction {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.transactions[address]; !ok {
		return []*parser.Transaction{}
	}

	return s.transactions[address]
}

func (s *MemoryStorage) SaveTransaction(address string, tx *parser.Transaction) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.transactions[address]; !ok {
		s.transactions[address] = []*parser.Transaction{}
	}

	s.transactions[address] = append(s.transactions[address], tx)
}
