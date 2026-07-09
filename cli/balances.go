package cli

import (
    "fmt"

    "github.com/spf13/cobra"
    "go_assesment/internal/blockchain"
    "go_assesment/internal/storage"
)

func init() {
    rootCli.AddCommand(balancesCmd)
}

var balancesCmd = &cobra.Command{
    Use:   "balances",
    Short: "Show account balances",
    RunE: func(cmd *cobra.Command, args []string) error {
        chainPath := dataDir + "/chain.json"
        var chain blockchain.Chain
        if err := storage.LoadJSON(chainPath, &chain); err != nil {
            return fmt.Errorf("could not load chain: %v", err)
        }
        ledger := blockchain.NewLedger()
        if _, err := ledger.RebuildfromChain(&chain); err != nil {
            return fmt.Errorf("rebuild error: %v", err)
        }
        fmt.Println("Balances:")
        for acct, bal := range ledger.Balances {
            fmt.Printf("  %s: %d\n", acct, bal)
        }
        return nil
    },
}
