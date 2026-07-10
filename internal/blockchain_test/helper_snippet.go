
package blockchain_test

import (
    

    "go_assesment/internal/blockchain"
)


func withDifficulty(d int, fn func()) {
    old := blockchain.GetDifficulty()
    blockchain.SetDifficulty(d)
    defer blockchain.SetDifficulty(old)
    fn()
}

func makeSimpleChain() blockchain.Chain {
    g := blockchain.NewGenesis()
    return blockchain.Chain{Blocks: []blockchain.Block{g}}
}
