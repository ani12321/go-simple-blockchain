package main

import (
	"fmt"
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
	}

	blockchain := &Blockchain{}
	blockchain.add(block)

	blockchain.startServer()
}
