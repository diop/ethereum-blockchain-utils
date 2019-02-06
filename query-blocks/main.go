package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	// Grab block header, then block number, then query on the actual block.
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	blockNumber := header.Number
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Block time: " + block.Time().String())
	fmt.Println("Block difficulty: " + block.Difficulty().String())
	fmt.Println("Block hash: " + block.Hash().String())
	fmt.Println("Number of TXs: " + strconv.Itoa(len(block.Transactions())))

}
