package cli

import (
	"fmt"
	"path/filepath"

	"os"
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


		var name string;

		if err := os.MkdirAll(dataDir, 0o755); err != nil {
			return fmt.Errorf("There is no directory to store pool or chain and failed to create dataDir: %w", err)
		}

			tx :=blockchain.Transaction{
			Sender: txFrom,
			Recipient: txTo,
			Amount: txAmt,
			Nonce: time.Now().UnixMilli(),
		}

		switch tx.Sender {
case "alice":name = "/alice.json"
	case "bob":name = "/bob.json"
case "kamal":name = "/kamal.json"
default:
	return fmt.Errorf("Sender not found : %s" ,tx.Sender)}


		senderPrivKey, err := blockchain.LoadPrivKey(filepath.Join(dataDir,"/keystore",name))

		if err != nil {
			return fmt.Errorf("failed to load private key: %v", err)
		}

		if err:= blockchain.SignTx(senderPrivKey, &tx);err != nil{
			return fmt.Errorf("Failed to sign tx :%v",err)
		}


		if !blockchain.VerifyTx(&tx){
			return fmt.Errorf("invalid signature")
		}

		chainPath := dataDir + "/chain.json"
		chainLock, err := storage.LockFile(chainPath, 5*time.Second)
		if err != nil {
			return err
		}
		defer chainLock.Unlock()
		var chain blockchain.Chain

		if err := storage.LoadJSON(chainPath, &chain); err != nil {
			gen := blockchain.NewGenesis()
			chain = blockchain.Chain{Blocks: []blockchain.Block{gen}}
			if err := storage.SaveJSON(chainPath, &chain); err != nil {
				return err
			}
		}

		ledger := blockchain.NewLedger()

		if _, err := ledger.RebuildfromChain(&chain); err != nil {
			return fmt.Errorf("Ledger rebuild error: %v", err)
		}

		if txFrom == "" || txTo == "" || txAmt <= 0 {
			return fmt.Errorf("provide --from, --to and --amt > 0")
		}

		bal := ledger.GetBalance(txFrom)
		if txAmt > bal {
			return fmt.Errorf("%s doesn't have enough balance: %d", txFrom, bal)
		}
		/*	else if txAmt > blockchain.NewLedger().GetBalance(txFrom) {
				return fmt.Errorf("%d doesn't have enough balance\n", blockchain.NewLedger().GetBalance(txFrom))
			}
		*/

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

		pool = append(pool, tx)
		if err := storage.SaveJSON(poolPath, pool); err != nil {
			return err
		}
		fmt.Println("Transaction added to the pool")
		return nil
	},
}
