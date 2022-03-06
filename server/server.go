package server

import (
	"context"
	"github.com/Kamran151199/gRPC-blockchain/proto"
)

type MyServer struct {
	Blockchain *Blockchain
}

func (s MyServer) AddBlock(ctx context.Context, request *proto.AddBlockRequest) (*proto.AddBlockResponse, error) {
	block := s.Blockchain.addBlock(request.Data)

	return &proto.AddBlockResponse{
		Hash: block.Hash,
	}, nil
}

func (s MyServer) GetBlockchain(context.Context, *proto.GetBlockchainRequest) (*proto.GetBlockchainResponse, error) {
	response := new(proto.GetBlockchainResponse)
	for _, block := range s.Blockchain.Blocks {
		response.Blocks = append(response.Blocks, &proto.Block{
			PrevBlockHash: block.PrevBlockHash,
			Hash:          block.Hash,
			Data:          block.Data,
		})
	}
	return response, nil
}
