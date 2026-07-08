package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

const difficulty = 3

type Transaction struct {
	Sender    string
	Recipient string
	Amount    int64
}

type Block struct {
	Index        uint64
	Timestamp    int64
	Txns         []Transaction
	Hash         string
	Nonce        uint64
	PreviousHash string
}

type Chain struct {
	Blocks []Block
}

type Ledger struct {
	Balances map[string]int64
}

func HashBlock(b *Block) string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "%d|%d|%s|%d|", b.Index, b.Timestamp, b.PreviousHash, b.Nonce)
	for _, t := range b.Txns {
		fmt.Fprintf(&buf, "%s>%s>%d;", t.Sender, t.Recipient, t.Amount)
	}
	sum := sha256.Sum256(buf.Bytes())
	return hex.EncodeToString(sum[:])

}

func NewGenesis() Block {
	genesis := Block{
		Index:     0,
		Timestamp: 0,
		Txns: []Transaction{
			{Sender: "coinbase", Recipient: "alice", Amount: 1000},
		},
		PreviousHash: strings.Repeat("0", 64),
		Nonce:        0,
	}
	genesis.Hash = HashBlock(&genesis)
	return genesis
}

func NewLedger() *Ledger {

	return &Ledger{Balances: make(map[string]int64)}
}

func (l *Ledger) GetBalance(account string) int64 {
	return l.Balances[account]
}

func (l *Ledger) ApplyTransaction(tx Transaction) error {
	if tx.Amount <= 0 {
		return fmt.Errorf("invalid amount")
	}

	if tx.Sender != "coinbase" {
		if l.GetBalance(tx.Sender) < tx.Amount {
			return fmt.Errorf("insufficient funds")
		}
		l.Balances[tx.Sender] -= tx.Amount
	}
	l.Balances[tx.Recipient] += tx.Amount
	return nil
}

// Rebuild balances from genesis to tip. Returns first offending block index and error if any
func (l *Ledger) RebuildfromChain(chain *Chain) (int, error) {
	l.Balances = make(map[string]int64)
	for i, blk := range chain.Blocks {
		for _, tx := range blk.Txns {
			if tx.Amount <= 0 {
				return i, fmt.Errorf("invalid amount in block %d", i)
			}
			if tx.Sender != "coinbase" {
				if l.GetBalance(tx.Sender) < tx.Amount {
					return i, fmt.Errorf("overspend by %s in block %d", tx.Sender, i)
				}
				l.Balances[tx.Sender] -= tx.Amount
			}
			l.Balances[tx.Recipient] += tx.Amount
		}
	}
	return -1, nil
}

func MineBlock(b *Block) {
	for {
		hash := HashBlock(b)
		if strings.HasPrefix(hash, strings.Repeat("0", difficulty)) {
			b.Hash = hash
			return
		}
		b.Nonce++
	}
}

func NewBlock(prev Block, txns []Transaction, timestamp int64) Block {
	block := Block{
		Index:        prev.Index + 1,
		Timestamp:    timestamp,
		Txns:         txns,
		PreviousHash: prev.Hash,
		Nonce:        0,
	}
	MineBlock(&block)
	return block
}

func ValidateChain(chain *Chain) (bool, error) {
	for i := 1; i < len(chain.Blocks); i++ {
		current := chain.Blocks[i]
		prev := chain.Blocks[i-1]

		if HashBlock(&current) != current.Hash {
			return false, fmt.Errorf("tampered hash at block %d", current.Index)
		}

		if current.PreviousHash != prev.Hash {
			return false, fmt.Errorf("Broken chain link at block %d", current.Index)
		}
	}

	ledger := NewLedger()
	if badIndex, err := ledger.RebuildfromChain(chain); err != nil {
		return false, fmt.Errorf("ledger inconsistency at block %d: %v", badIndex, err)
	}

	return true, nil
}



func main() {
	genesis := NewGenesis()
	chain := Chain{Blocks: []Block{genesis}}

	txns := []Transaction{
		{Sender: "alice", Recipient: "bob", Amount: 50},
		{Sender: "bob", Recipient: "alice", Amount: 24},
		{Sender: "alice", Recipient: "kamal", Amount: 500},
	}
	newBlock := NewBlock(genesis, txns, 1)
	chain.Blocks = append(chain.Blocks, newBlock)

	fmt.Println("Genesis Hash:", genesis.Hash)
	fmt.Println("Mined Block Hash:", newBlock.Hash)
	fmt.Println("Nonce:", newBlock.Nonce)

	ok, error := ValidateChain(&chain)
	if !ok {
		fmt.Println("Chain invalid:", error)
	} else {
		fmt.Println("Chain is valid.")

		ledger := NewLedger()
		ledger.RebuildfromChain(&chain)
		fmt.Println("Balances:", ledger.Balances)
	}

}
