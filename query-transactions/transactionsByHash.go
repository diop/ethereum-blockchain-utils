package main()

func transactionsByHashHandler() {
	client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	txHash := common.HexToHash("0x3cef650db672f33d10f810073b2280f801fecf91bbd100e09d7d26f2eb703031")
}