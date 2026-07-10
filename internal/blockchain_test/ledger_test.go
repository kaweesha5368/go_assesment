package blockchain_test

import (
    "testing"

    "go_assesment/internal/blockchain"
)

func TestRebuildfromChain_Overspend(t *testing.T) {
    g := blockchain.NewGenesis()
    // create a block where alice overspends
    bad := blockchain.Block{
        Index: 1, Timestamp: 1,
        Txns: []blockchain.Transaction{
            {Sender:"alice", Recipient:"bob", Amount: 1001},
        },
        PreviousHash: g.Hash,
        Hash: "", Nonce: 0,
    }
    chain := blockchain.Chain{Blocks: []blockchain.Block{g, bad}}
    l := blockchain.NewLedger()
    idx, err := l.RebuildfromChain(&chain)
    if err == nil {
        t.Fatal("expected overspend error")
    }
    if idx != 1 {
        t.Fatalf("expected bad index 1 got %d", idx)
    }
}
