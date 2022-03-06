package main

import (
	"context"
	"flag"
	"github.com/Kamran151199/gRPC-blockchain/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

var client proto.BlockchainClient

func main() {
	addFlag := flag.Bool("add", false, "add new block")
	listFlag := flag.Bool("list", false, "get the blockchain")
	flag.Parse()
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
	}

	client = proto.NewBlockchainClient(conn)

	if *addFlag {
		addBlock()
	}
	if *listFlag {
		getBlockchain()
	}
}

func addBlock() {
	block, err := client.AddBlock(context.Background(), &proto.AddBlockRequest{
		Data: time.Now().String(),
	})
	if err != nil {
		log.Fatalf("unable to create block: %v", err)
	}
	log.Printf("New block hash: %v", block.Hash)
}

func getBlockchain() {
	blockchain, err := client.GetBlockchain(context.Background(), &proto.GetBlockchainRequest{})
	if err != nil {
		log.Fatalf("unable to fetch blockchain: %v", err)
	}
	for index, block := range blockchain.Blocks {
		log.Printf("Block #%v hash: %s, previous hash: %s, data: %s\n",
			index, block.Hash, block.PrevBlockHash, block.Data)
	}
}
