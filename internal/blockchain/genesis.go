package blockchain

import "strings"

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
