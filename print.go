
package cli

import (
    "fmt"

    "github.com/spf13/cobra"
    "go_assesment/internal/blockchain"
    "go_assesment/internal/storage"
)

func init() {
    rootCli.AddCommand(printCli)
}

var printCli = &cobra.Command{
    Use:   "print",
    Short: "Print the chain in readable form",
    RunE: func(cmd *cobra.Command, args []string) error {
        chainPath := dataDir + "/chain.json"
        var chain blockchain.Chain
        if err := storage.LoadJSON(chainPath, &chain); err != nil {
            return fmt.Errorf("could not load chain: %v", err)
        }
        for _, b := range chain.Blocks {
            fmt.Printf("Index: %d  Timestamp: %d\n", b.Index, b.Timestamp)
            fmt.Printf("PrevHash: %s\nHash: %s\nNonce: %d\n", b.PreviousHash, b.Hash, b.Nonce)
            fmt.Println("Transactions:")
            for _, t := range b.Txns {
                fmt.Printf("  %s -> %s : %d\n", t.Sender, t.Recipient, t.Amount)
            }
            fmt.Println("----")
        }
        return nil
    },
}