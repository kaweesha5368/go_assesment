package blockchain

import "fmt"

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
