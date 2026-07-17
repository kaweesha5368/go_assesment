package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	//"strconv"

	//"github.com/golang/protobuf/ptypes/empty"
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

/*
func HashBlock(b *Block) string {
    var leaves [][]byte
    for _, t := range b.Txns {
        leaves=append(leaves, TxLeafHash(&t))
    }
    mr := MerkleRoot(leaves)
    b.MerkleRoot = hex.EncodeToString(mr)

    var buf bytes.Buffer
    fmt.Fprintf(&buf, "%d|%d|%d|%s|%d|%s", b.Index,b.Timestamp, b.PreviousHash,b.Nonce, b.MerkleRoot)
    sum := sha256.Sum256(buf.Bytes())
    return hex.EncodeToString(sum[:])
}
*/
/*
func TxHash(tx Transaction)[]byte{
    b:=bytes.Join([][]byte{
        []byte(tx.Sender),
        []byte(tx.Recipient),
        []byte(strconv.FormatInt(tx.Amount,10)),
        []byte(strconv.FormatInt(tx.Nonce,10)),
    },[]byte("|"))
    h:=sha256.Sum256(b)
    return h[:]
}
*/

/*
func MerkleRoot(hashes [][]byte) []byte{
    if len(hashes) == 0 {
        empty := sha256.Sum256([]byte{})
        return empty[:]
    }

    for len(hashes) > 1 {
        if len(hashes)%2 == 1{
            hashes = append(hashes, hashes[len(hashes)-1])
        }
        var next [][]byte
        for i := 0; i<len(hashes); i+=2{
            combined := append(hashes[i], hashes[i+1]...)
            sum := sha256.Sum256(combined)
            next = append(next, sum[:])
        }
        hashes = next
    }
    return hashes[0]
}*/