package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"go_assesment/internal/blockchain"
	"go_assesment/internal/storage"
)

var miner string
var reward int64 = 50

func init() {
	mineCli.Flags().StringVar(&miner, "miner", "miner", "miner account name")
	mineCli.Flags().Int64Var(&reward, "reward", 50, "coinbase reward")
	rootCli.AddCommand(mineCli)
}

var mineCli = &cobra.Command{
	Use:   "mine",
	Short: "Mine a block from pending pool",
	RunE: func(cmd *cobra.Command, args []string) error {
		chainPath := dataDir + "/chain.json"
		poolPath := dataDir + "/pool.json"

		// lock chain and pool
		chainLock, err := storage.LockFile(chainPath, 5*time.Second)
		if err != nil {
			return err
		}
		defer chainLock.Unlock()

		poolLock, err := storage.LockFile(poolPath, 5*time.Second)
		if err != nil {
			return err
		}
		defer poolLock.Unlock()

		var chain blockchain.Chain
		if err := storage.LoadJSON(chainPath, &chain); err != nil {
			// initialize with genesis if missing
			gen := blockchain.NewGenesis()
			chain = blockchain.Chain{Blocks: []blockchain.Block{gen}}
			if err := storage.SaveJSON(chainPath, &chain); err != nil {
				return err
			}
		}

		var pool []blockchain.Transaction
		if err := storage.LoadJSON(poolPath, &pool); err != nil {
			pool = []blockchain.Transaction{}
		}
		if len(pool) == 0 {
			return fmt.Errorf("no transactions in pool to mine")
		}

		// create coinbase tx
		coinbase := blockchain.Transaction{Sender: "coinbase", Recipient: miner, Amount: reward}
		txs := append([]blockchain.Transaction{coinbase}, pool...)

		prev := chain.Blocks[len(chain.Blocks)-1]
		
		newBlk := blockchain.NewBlock(prev, txs, time.Now().Unix())
		
		chain.Blocks = append(chain.Blocks, newBlk)

		// validate before saving
		if ok, err := blockchain.ValidateChain(&chain); !ok {
			return fmt.Errorf("validation failed after mining: %v", err)
		}

		if err := storage.SaveJSON(chainPath, &chain); err != nil {
			return err
		}

		// clear pool
		if err := storage.SaveJSON(poolPath, []blockchain.Transaction{}); err != nil {
			return err
		}

		fmt.Printf("Mined block %d \n", newBlk.Index)
		fmt.Printf("Hash=%s \n", newBlk.Hash)
		fmt.Printf("Nonce=%d \n", newBlk.Nonce)

		return nil
	},
}
