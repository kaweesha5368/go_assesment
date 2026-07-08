package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"

)

type Transaction struct{
	Sender string 	
	Recipient string 
	Amount int64 
}


type Block struct {
	Index uint64 
	Timestamp int64 
	Txns []Transaction
	Hash string 
	Nonce uint64 
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
	fmt.Fprintf(&buf, "%d|%d|%s|%d|", b.Index, b.Timestamp, b.PreviousHash,b.Nonce)
	for _, t := range b.Txns{
		fmt.Fprintf(&buf, "%s>%s>%d;", t.Sender, t.Recipient, t.Amount)
	}
	sum := sha256.Sum256(buf.Bytes())
	return hex.EncodeToString(sum[:])
	
}

func NewGenesis() Block{
	genesis :=Block{
		Index: 0,
		Timestamp :0,
		Txns: []Transaction{
			{Sender:"coinbase", Recipient: "alice", Amount:1000},
		},
		PreviousHash: strings.Repeat("0", 64),
		Nonce: 0,
			}
			genesis.Hash = HashBlock(&genesis)
			return genesis
}


func NewLedger() *Ledger {

	return &Ledger{Balances: make(map[string]int64)}
}

func (l *Ledger) GetBalance(account string) int64{
	return l.Balances[account]
}

func (l *Ledger) ApplyTransaction(tx Transaction) error {
	if tx.Amount <= 0 {
		return fmt.Errorf("invalid amount")
	}

	if tx.Sender !="coinbase" {
		if l.GetBalance(tx.Sender) < tx.Amount{
			return fmt.Errorf("insufficient funds")
		}
		l.Balances[tx.Sender] -= tx.Amount
	}
	l.Balances[tx.Recipient] += tx.Amount
	return nil
}

//Rebuild balances from genesis to tip. Returns first offending block index and error if any
func (l *Ledger) RebuildfromChain(chain *Chain) (int, error) {
	l.Balances = make(map[string]int64)
	for i, blk := range chain.Blocks {
		for _, tx := range blk.Txns {
			if tx.Amount <= 0 {
				return i, fmt.Errorf("invalid amount in block %d", i)
			}
			if tx.Sender != "coinbase" {
				if l.GetBalance(tx.Sender) < tx.Amount {
					return i, fmt.Errorf("overspend by %s in block %d",tx.Sender, i)
				}
				l.Balances[tx.Sender] -=tx.Amount
			}
			l.Balances[tx.Recipient] += tx.Amount
		}
	}
	return -1, nil
}

func main() {
	fmt.Println("HelloWorld")
}
