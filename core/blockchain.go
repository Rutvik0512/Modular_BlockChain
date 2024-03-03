package core

import (
	"fmt"
	"sync"
)

type BlockChain struct {
	store     Storage
	lock      sync.RWMutex
	headers   []*Header
	validator Validator
}

func NewBlockChain(genesis *Block) (*BlockChain, error) {
	bc := &BlockChain{
		headers: []*Header{},
		store:   NewMemoryStore(),
	}
	bc.validator = NewBlockValidator(bc)
	err := bc.addBlockWithoutValidation(genesis)

	return bc, err
}

func (bc *BlockChain) SetValidator(v Validator) {
	bc.validator = v
}

func (bc *BlockChain) AddBlock(b *Block) error {
	if err := bc.validator.ValidateBlock(b); err != nil {
		return err
	}

	return bc.addBlockWithoutValidation(b)
}

func (bc *BlockChain) GetHeader(height uint32) (*Header, error) {
	if height > bc.Height() {
		return nil, fmt.Errorf("given height (%d) too high", height)
	}

	bc.lock.Lock()
	defer bc.lock.Unlock()

	return bc.headers[height], nil
}

func (bc *BlockChain) HasBlock(height uint32) bool {
	return height <= bc.Height()
}

func (bc *BlockChain) Height() uint32 {
	bc.lock.RLock()
	defer bc.lock.RUnlock()

	return uint32(len(bc.headers) - 1)
}

func (bc *BlockChain) addBlockWithoutValidation(b *Block) error {
	bc.lock.Lock()

	bc.headers = append(bc.headers, b.Header)
	bc.lock.Unlock()

	// logrus.WithFields(logrus.Fields{
	// 	"height": b.Height,
	// 	"hash":   b.Hash(BlockHasher{}),
	// }).Info("adding new Block")
	return bc.store.Put(b)
}
