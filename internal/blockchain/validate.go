package blockchain

import (
	"fmt"
	"strings"
)

func ValidateChain(chain *Chain) (bool, error) {
    for i := 1; i < len(chain.Blocks); i++ {
        current := chain.Blocks[i]
        prev := chain.Blocks[i-1]

        if HashBlock(&current) != current.Hash {
            return false, fmt.Errorf("tampered hash at block %d", current.Index)
        }
        targetPrefix :=strings.Repeat("0",GetDifficulty())
        if !strings.HasPrefix(current.Hash,targetPrefix){
            return false, fmt.Errorf("block %d does not meet difficulty target",current.Index)
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


