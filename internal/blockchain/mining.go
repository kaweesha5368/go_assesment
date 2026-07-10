package blockchain

import "strings"
import "fmt"
import "time"

var Difficulty = 3

func SetDifficulty(d int) { Difficulty = d }
func GetDifficulty() int  { return Difficulty }

func MineBlock(b *Block) (int, time.Duration) {
	start := time.Now()
	attempts := 0
	for {
		attempts++
		hash := HashBlock(b)
		if strings.HasPrefix(hash, strings.Repeat("0", GetDifficulty())) {
			b.Hash = hash

			return attempts, time.Since(start)
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
	attempts, elapsed := MineBlock(&block)
	fmt.Printf("Block mined in %s after %d attempts\n", elapsed, attempts)
	return block
}
