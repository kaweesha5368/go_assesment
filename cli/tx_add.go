
package cli

import (
	"fmt"
	"time"

	"go_assesment/internal/blockchain"
	"go_assesment/internal/storage"

	"github.com/spf13/cobra"
)

var txFrom, txTo string
var txAmt int64

func init() {
    txAddCli.Flags().StringVar(&txFrom, "from", "", "sender")
    txAddCli.Flags().StringVar(&txTo, "to", "", "recipient")
    txAddCli.Flags().Int64Var(&txAmt, "amt", 0, "amount")
    rootCli.AddCommand(txAddCli)
}

var txAddCli = &cobra.Command{
    Use:   "tx add",
    Short: "Add transaction to pending pool",
    RunE: func(cmd *cobra.Command, args []string) error {
        if txFrom == "" || txTo == "" || txAmt <= 0 {
            return fmt.Errorf("provide --from, --to and --amt > 0")
        }
        poolPath := dataDir + "/pool.json"
        // lock pool file
        lock, err := storage.LockFile(poolPath, 5*time.Second)
        if err != nil {
            return err
        }
        defer lock.Unlock()

        var pool []blockchain.Transaction
        if err := storage.LoadJSON(poolPath, &pool); err != nil {
            // if file missing, start with empty pool
            pool = []blockchain.Transaction{}
        }
        pool = append(pool, blockchain.Transaction{Sender: txFrom, Recipient: txTo, Amount: txAmt})
        if err := storage.SaveJSON(poolPath, pool); err != nil {
            return err
        }
        fmt.Println("Transaction added to pool.")
        return nil
    },
}