package parser

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
	if !ep.subscribed[address] {
		return nil
	}
	return ep.transactions[address]
}
