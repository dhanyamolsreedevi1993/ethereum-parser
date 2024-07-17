package storage

import "github.com/dhanyamolsreedevi1993/ethereum-parser/parser"

// Storage defines the interface for transaction storage logic.
type Storage interface {
	GetCurrentBlock() int // Not implemented in memory storage
	GetTransactions(address string) []*parser.Transaction
	SaveTransaction(address string, tx *parser.Transaction)
}
