package blockchain_test

import (
	"testing"
	"time"

	"go_assesment/internal/blockchain"
)

func TestNewBlockAndMineLowDifficulty(t *testing.T) {
	withDifficulty := func(d int, fn func()) {
		old := blockchain.Difficulty
		blockchain.Difficulty = d
		defer func() { blockchain.Difficulty = old }()
		fn()
	}

	withDifficulty(1, func() {
		prev := blockchain.NewGenesis()
		txs := []blockchain.Transaction{
			{Sender: "coinbase", Recipient: "alice", Amount: 10},
		}
		blk := blockchain.NewBlock(prev, txs, time.Now().Unix())
        blk.Nonce++
		if blk.Hash == "" {
			t.Fatal("expected non-empty hash")
		}
		if blockchain.HashBlock(&blk) != blk.Hash {
			t.Fatalf("hash mismatch")
		}
	})
}
