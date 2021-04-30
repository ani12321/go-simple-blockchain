package main

import (
	"fmt"
	"net"
	"sync"
)

type Blockchain struct {
	Data       []Block
	restServer *RestServer
	tcpServer  *TcpServer
	sockets    []*net.Conn
	difficulty int
}

func (b *Blockchain) getLatestBlock() *Block {
	var l = len(b.Data)
	if l == 0 {
		return nil
	}
	return &b.Data[l-1]
}

func (b *Blockchain) add(block *Block) {
	latestBlock := b.getLatestBlock()
	if latestBlock != nil {
		block.PreviousHash = latestBlock.Hash
	}
	block.Hash = block.computeHash()
	fmt.Println("Applying proof of work")
	block.proofOfWork(b.difficulty)
	fmt.Println("Proof of work complete")
	b.Data = append(b.Data, *block)
}

func (b *Blockchain) checkValidity() bool {
	for i := 1; i < len(b.Data); i++ {
		currentBlock := b.Data[i]
		previousBlock := b.Data[i-1]

		if currentBlock.Hash != currentBlock.computeHash() {
			return false
		}
		if currentBlock.PreviousHash != previousBlock.Hash {
			return false
		}
	}
	return true
}

func (b *Blockchain) startServer() {
	b.restServer = &RestServer{port: 8089}

	var wg sync.WaitGroup
	// wg.Add(1)
	// go b.restServer.start(b, &wg)
	wg.Add(1)
	go b.tcpServer.start(b, &wg)
	wg.Wait()
}

func (b *Blockchain) removePeer(socket *net.Conn) {
	var index = -1
	for i, n := range b.sockets {
		if socket == n {
			index = i
		}
	}

	if index != -1 {
		b.sockets = append(b.sockets[:index], b.sockets[index+1:]...)
	}
	fmt.Println((*socket).RemoteAddr().String() + " peer disconnected")
}
