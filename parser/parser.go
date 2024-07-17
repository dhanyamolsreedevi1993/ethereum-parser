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
}

func NewEthereumParser() *EthereumParser {
	return &EthereumParser{
		subscribed: make(map[string]bool),
	}
}

func (ep *EthereumParser) GetCurrentBlock() int {
	return ep.currentBlock
}

func (ep *EthereumParser) Subscribe(address string) bool {
	ep.subscribed[address] = true
	return true
}

func (ep *EthereumParser) GetTransactions(address string) []*Transaction {
	return []*Transaction{}
}
