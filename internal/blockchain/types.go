package blockchain

type Transaction struct {
    Sender    string `json:"sender"`
    Recipient string `json:"recipient"`
    Amount    int64  `json:"amount"`
}

type Block struct {
    Index        uint64        `json:"index"`
    Timestamp    int64         `json:"timestamp"`
    Txns         []Transaction `json:"txns"`
    Hash         string        `json:"hash"`
    Nonce        uint64        `json:"nonce"`
    PreviousHash string        `json:"previous_hash"`
}

type Chain struct {
    Blocks []Block `json:"blocks"`
}

type Ledger struct {
    Balances map[string]int64 `json:"balances"`
}
