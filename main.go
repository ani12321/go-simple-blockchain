package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Initializing")

	block := &Block{
		Index:        0,
		PreviousHash: "0",
		Timestamp:    time.Unix(1465154705, 0),
		Data:         "my genesis block",
		Hash:         "816534932c2b7154836da6afc367695e6337db8a921823784c14378abed4f7d7",
		Nonce:        0,
	}

	block2 := &Block{
		Index:     1,
		Timestamp: time.Unix(1465154705, 0),
		Data:      "my second block",
		Nonce:     0,
	}

	blockchain := &Blockchain{
		difficulty: 4,
	}
	blockchain.add(block)
	blockchain.add(block2)

	validChain := blockchain.checkValidity()
	if !validChain {
		fmt.Println("Blockchain is not valid")
		os.Exit(1)
	}
	fmt.Println("Blockchain validity successful")
	fmt.Println(blockchain)

	blockchain.startServer()
}
