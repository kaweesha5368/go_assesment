package cli

import (
    "fmt"

    "github.com/spf13/cobra"
    "go_assesment/internal/blockchain"
    "go_assesment/internal/storage"
)

func init() {
    rootCli.AddCommand(validateCli)
}

var validateCli = &cobra.Command{
    Use:   "validate",
    Short: "Validate the chain",
    RunE: func(cmd *cobra.Command, args []string) error {
        chainPath := dataDir + "/chain.json"
        var chain blockchain.Chain
        if err := storage.LoadJSON(chainPath, &chain); err != nil {
            return fmt.Errorf("could not load chain: %v", err)
        }
        ok, err := blockchain.ValidateChain(&chain)
        if !ok {
            fmt.Println("Chain invalid:", err)
        } else {
            fmt.Println("Chain is valid.")
        }
        return nil
    },
}
