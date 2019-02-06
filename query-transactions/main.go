package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	blockNumber := big.NewInt(6222222)

	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	for _, tx := range block.Transactions() {
		fmt.Printf("Tx Hash: %s\n", tx.Hash().Hex())
		fmt.Printf("Tx Value: %s\n", tx.Value().String())
		fmt.Printf("Tx Gas: %s\n", tx.Gas())
		fmt.Printf("Tx Gas Price: %s\n", tx.GasPrice().Uint64())
		fmt.Printf("Tx Nonce: %s\n", tx.Nonce())
		fmt.Printf("Tx Nonce: %s\n", tx.Data())
		fmt.Printf("Tx Data: %s\n", tx.To().Hex())

		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Receipt status: %d\n", receipt.Status)
		fmt.Println("---")
	}

	// We can also query transactions by index.
}
