package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	conn, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatalf("Something went wrong connecting to the Ethereum node.", err)
	}

	ctx := context.Background()
	tx, pending, _ := conn.TransactionByHash(ctx, common.HexToHash("0xaf627d0f5912595dc7aa9897190441086692a40410384398b30d95c389bc8dd7"))
	if !pending {
		fmt.Println(tx)
	}
}
