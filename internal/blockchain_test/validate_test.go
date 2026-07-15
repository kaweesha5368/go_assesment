package blockchain_test

import (
	//"fmt"
	"fmt"
	"testing"
	"time"

	"go_assesment/internal/blockchain"
	"strings"
)

func TestValidateChain_TamperedHash(t *testing.T) {
	withDifficulty := func(d int, fn func()) {
		old := blockchain.Difficulty
		blockchain.Difficulty = d
		defer func() { blockchain.Difficulty = old }()
		fn()
	}

	withDifficulty(2, func() {
		g := blockchain.NewGenesis()
		blk := blockchain.NewBlock(g, []blockchain.Transaction{
			{Sender: "coinbase", Recipient: "alice", Amount: 1002},
		}, time.Now().Unix())
		// tamper
		fmt.Println(blk.Txns[0].Amount)
		blk.Txns[0].Amount = 1003
		//when blk.Txns is 1002 test fails as it can't find a mismatch between the original and modified amounts
		chain := blockchain.Chain{Blocks: []blockchain.Block{g, blk}}
		ok, err := blockchain.ValidateChain(&chain)
        
		if ok {
			t.Fatalf("expected ValidateChain to return ok=false for tampered chain")
		}
		if err == nil {
			t.Fatalf("expected ValidateChain to return an error for tampered chain")
		}
		// optional: assert on error text
		if !strings.Contains(err.Error(), "tampered") && !strings.Contains(err.Error(), "ledger inconsistency") {
			t.Fatalf("unexpected validation error: %v", err)
		}
	})
}
