package blockchain_test

import (
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
		blk.Txns[2].Amount = 1003
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
