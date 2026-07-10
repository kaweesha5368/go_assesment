
package blockchain

import "strings"

var Difficulty = 3

func SetDifficulty(d int) {Difficulty = d}
func GetDifficulty() int { return Difficulty}

func MineBlock(b *Block) {
    for {
        hash := HashBlock(b)
        if strings.HasPrefix(hash, strings.Repeat("0", GetDifficulty())) {
            b.Hash = hash
            return
        }
        b.Nonce++
    }
}

func NewBlock(prev Block, txns []Transaction, timestamp int64) Block {
    block := Block{
        Index:        prev.Index + 1,
        Timestamp:    timestamp,
        Txns:         txns,
        PreviousHash: prev.Hash,
        Nonce:        0,
    }
    MineBlock(&block)
    return block
}