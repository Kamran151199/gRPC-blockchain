package server

func NewBlock(data string, previousBlockHash string) *Block {
	block := &Block{
		Data:          data,
		PrevBlockHash: previousBlockHash,
	}
	block.setHash()
	return block
}

func CreateBlockchain() *Blockchain {
	return &Blockchain{[]*Block{CreateGenesis()}}
}

func CreateGenesis() *Block {
	block := NewBlock("Kamran is the best", "")
	return block
}
