package blockchain

import (
    "bytes"
    "crypto/sha256"
    "encoding/hex"
    "fmt"
)

func HashBlock(b *Block) string {
    var buf bytes.Buffer
    fmt.Fprintf(&buf, "%d|%d|%s|%d|", b.Index, b.Timestamp, b.PreviousHash, b.Nonce)
    for _, t := range b.Txns {
        fmt.Fprintf(&buf, "%s>%s>%d;", t.Sender, t.Recipient, t.Amount)
    }
    sum := sha256.Sum256(buf.Bytes())
    return hex.EncodeToString(sum[:])
}
