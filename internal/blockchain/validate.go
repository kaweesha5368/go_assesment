package blockchain

import "fmt"

func ValidateChain(chain *Chain) (bool, error) {
    for i := 1; i < len(chain.Blocks); i++ {
        current := chain.Blocks[i]
        prev := chain.Blocks[i-1]

        if HashBlock(&current) != current.Hash {
            return false, fmt.Errorf("tampered hash at block %d", current.Index)
        }
        if current.PreviousHash != prev.Hash {
            return false, fmt.Errorf("broken chain link at block %d", current.Index)
        }
    }

    ledger := NewLedger()
    if badIndex, err := ledger.RebuildfromChain(chain); err != nil {
        return false, fmt.Errorf("ledger inconsistency at block %d: %v", badIndex, err)
    }
    return true, nil
}
