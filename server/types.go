package server

import (
	"crypto/sha256"
	"encoding/hex"
)

type Block struct {
	Hash          string
	PrevBlockHash string
	Data          string
}

func (b *Block) setHash() {
	hash := sha256.Sum256([]byte(b.PrevBlockHash + b.Data))
	b.Hash = hex.EncodeToString(hash[:])
}

type Blockchain struct {
	Blocks []*Block
}

func (bch *Blockchain) addBlock(data string) *Block {
	prevBlock := bch.Blocks[(len(bch.Blocks) - 1)]
	newBlock := NewBlock(data, prevBlock.Hash)
	bch.Blocks = append(bch.Blocks, newBlock)
	return newBlock
}
