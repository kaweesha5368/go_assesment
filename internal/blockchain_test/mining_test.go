
package blockchain_test

import (
	"testing"
	"time"
	"fmt"
	"go_assesment/internal/blockchain"
)

func TestNewBlockAndMineLowDifficulty(t *testing.T) {
	withDifficulty := func(d int, fn func()) {
		old := blockchain.Difficulty
		blockchain.Difficulty = d
		defer func() { blockchain.Difficulty = old }()
		fn()
		fmt.Printf("Difficulty = %d \n", d)
	}

	withDifficulty(3, func()  {
		prev := blockchain.NewGenesis()
		txs := []blockchain.Transaction{
			{Sender: "coinbase", Recipient: "alice", Amount: 10},
		}
		blk := blockchain.NewBlock(prev, txs, time.Now().Unix())
        //blk.Nonce++  <-- to make the test fail, when nonce is changed
		if blk.Hash == "" {
			t.Fatal("expected non-empty hash")
		}
		if blockchain.HashBlock(&blk) != blk.Hash {
			t.Fatalf("hash mismatch")
		}

		return
	})
}