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
        var transactions []blockchain.Transaction
        if err := storage.LoadJSON(poolPath, &transactions); err !=nil{
            return fmt.Errorf("could not load pool: %v", err)
        }
      
      for _, t := range transactions {
        fmt.Printf(" %s -> %s : %d\n", t.Sender, t.Recipient, t.Amount)
      }
            fmt.Println("----")

            return nil
    },

}