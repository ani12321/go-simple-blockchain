package main

import (
	"sync"
	"time"
)

type Block struct {
	Index        int
	PreviousHash string
	Timestamp    time.Time
	Data         string
	Hash         string
}

type Blockchain struct {
	Data       []Block
	restServer *RestServer
	tcpServer  *TcpServer
}

func (b *Blockchain) getLatestBlock() *Block {
	var l = len(b.Data)
	return &b.Data[l-1]
}

func (b *Blockchain) add(block *Block) {
	b.Data = append(b.Data, *block)
}

func (b *Blockchain) startServer() {
	b.restServer = &RestServer{port: 8089}

	var wg sync.WaitGroup

	wg.Add(1)
	go b.restServer.start(b, &wg)
	wg.Add(1)
	go b.tcpServer.start(b, &wg)

	wg.Wait()
}
