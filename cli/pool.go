package cli

import (
   "fmt"

    "github.com/spf13/cobra"
    "go_assesment/internal/blockchain"
    "go_assesment/internal/storage"
)

func init(){
    rootCli.AddCommand(poolCli)
}

var poolCli = &cobra.Command{
    Use : "pool",
    Short: "Print the transaction waiting in the pool",
    RunE: func(cmd *cobra.Command, args []string) error {
        poolPath := dataDir + "/pool.json"
        var transaction blockchain.Transaction
        if err := storage.LoadJSON(poolPath, &transaction);err !=nil{
            return fmt.Errorf("could not load pool: %v", err)
        }
        for _, b :=range pool
    }
}