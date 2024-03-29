package core

import "fmt"

type Validator interface {
	ValidateBlock(*Block) error
}

type BlockValidator struct {
	bc *BlockChain
}

func NewBlockValidator(bc *BlockChain) *BlockValidator {
	return &BlockValidator{
		bc: bc,
	}
}

func (v *BlockValidator) ValidateBlock(b *Block) error {
	if v.bc.HasBlock(b.Height) {
		return fmt.Errorf("chain already has block (%d) with hash (%s)", b.Height, b.Hash(BlockHasher{}))
	}

	if b.Height != v.bc.Height()+1 {
		return fmt.Errorf("block too high")
	}

	if err := b.Verify(); err != nil {
		return err
	}

	prevHeader, err := v.bc.GetHeader(b.Height - 1)
	hash := BlockHasher{}.Hash(prevHeader)

	if err != nil {
		return err
	}

	if hash != b.PrevBlockHash {
		return fmt.Errorf("the hash of the previous block (%d) is invalid", b.PrevBlockHash)
	}

	return nil
}
