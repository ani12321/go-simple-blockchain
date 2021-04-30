package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// type BlockData struct {
// 	Sender string `json:"sender"`
// 	Receiver string `json:"receiver"`
// }

type Block struct {
	Index        int
	PreviousHash string
	Timestamp    time.Time
	Data         string
	Hash         string
	Nonce        int
}

func (b *Block) computeHash() string {
	str := fmt.Sprintf("%d%s%d%s%d", b.Index, b.PreviousHash, b.Timestamp.Unix(), b.Data, b.Nonce)
	hash := sha256.Sum256([]byte(str))
	return hex.EncodeToString(hash[:])
}

func (b *Block) proofOfWork(difficulty int) {
	for {
		c := []byte(b.Hash[0:difficulty])
		d := make([]byte, difficulty)
		for i := 0; i < difficulty; i++ {
			d[i] = byte('0')
		}
		if !bytes.Equal(c, d) {
			b.Nonce = b.Nonce + 1
			b.Hash = b.computeHash()
		} else {
			return
		}
	}
}
