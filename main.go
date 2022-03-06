package main

import (
	"github.com/Kamran151199/gRPC-blockchain/proto"
	"github.com/Kamran151199/gRPC-blockchain/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	// listener for our gRPC server
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalf("unable to listen on port 8080 %v", err)
	}

	// gRPC server to be passed as param to registration.
	grpcServer := grpc.NewServer()
	proto.RegisterBlockchainServer(grpcServer, server.MyServer{Blockchain: server.CreateBlockchain()})

	// start server
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
