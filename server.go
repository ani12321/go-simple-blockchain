package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"sync"
)

type RestServer struct {
	port uint16
}

func (s *RestServer) start(blockchain *Blockchain, wg *sync.WaitGroup) {

	hello := func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "hello\n", blockchain)
	}

	getBlocks := func(w http.ResponseWriter, req *http.Request) {

		data, err := json.Marshal(blockchain.Data)
		if err == nil {
			fmt.Fprintf(w, "Error")
			return
		}
		fmt.Fprintf(w, string(data))
	}

	mineBlocks := func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "POST" {
			return
		}

	}

	getPeers := func(w http.ResponseWriter, req *http.Request) {

	}

	addPeer := func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "POST" {
			return
		}
	}

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/blocks", getBlocks)
	http.HandleFunc("/mine", mineBlocks)
	http.HandleFunc("/addPeer", addPeer)
	http.HandleFunc("/peers", getPeers)

	fmt.Println("server starting at", s.port)
	http.ListenAndServe(":"+fmt.Sprint(s.port), nil)
}

type TcpServer struct {
}

func handlePeerConnection(blockchain *Blockchain, client *net.Conn) {
	conn := *client
	blockchain.sockets = append(blockchain.sockets, client)
	fmt.Println("Peer connected", conn.RemoteAddr())
	defer conn.Close()
	defer blockchain.removePeer(client)
	buf := make([]byte, 1024)
	for {
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			break
		}
		fmt.Printf(conn.RemoteAddr().String() + ":" + string(buf))
		conn.Write([]byte("Message received."))
		buf = make([]byte, 1024)
	}
}

func (t *TcpServer) start(blockchain *Blockchain, wg *sync.WaitGroup) {
	conn, err := net.Listen("tcp", "localhost:8898")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer conn.Close()
	defer wg.Done()
	fmt.Println("TCP connection listening on", conn.Addr())
	for {
		// Listen for an incoming connection.
		client, err := conn.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		go handlePeerConnection(blockchain, &client)
	}
}
