package core

import (
	"main/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func newBlockChainWithGenesis(t *testing.T) *BlockChain {
	bc, err := NewBlockChain(randomBlock(0, types.Hash{}))
	assert.Nil(t, err)

	return bc
}
func TestNewBlockChain(t *testing.T) {

	bc := newBlockChainWithGenesis(t)
	assert.NotNil(t, bc.validator)
	assert.Equal(t, bc.Height(), uint32(0))
}

func TestAddBlock(t *testing.T) {

	bc := newBlockChainWithGenesis(t)

	lenBlock := 1000
	for i := 0; i < lenBlock; i++ {
		b := randomBlockWithValidator(t, uint32(i+1), getPrevBlockHash(t, bc, uint32(i+1)))
		assert.Nil(t, bc.AddBlock(b))
	}

	assert.Equal(t, bc.Height(), uint32(lenBlock))
	assert.NotNil(t, bc.AddBlock(randomBlock(89, types.Hash{})))

}

func TestHasBlock(t *testing.T) {
	bc := newBlockChainWithGenesis(t)
	assert.True(t, bc.HasBlock(0))
}

func TestGetHeader(t *testing.T) {
	bc := newBlockChainWithGenesis(t)
	lenBlock := 1000

	for i := 0; i < lenBlock; i++ {
		b := randomBlockWithValidator(t, uint32(i+1), getPrevBlockHash(t, bc, uint32(i+1)))
		assert.Nil(t, bc.AddBlock(b))
		header, err := bc.GetHeader(uint32(i + 1))
		assert.Nil(t, err)
		assert.Equal(t, header, b.Header)
	}

}

func TestAddBlockHigh(t *testing.T) {
	bc := newBlockChainWithGenesis(t)
	assert.Nil(t, bc.AddBlock(randomBlockWithValidator(t, 1, getPrevBlockHash(t, bc, uint32(1)))))
	assert.NotNil(t, bc.AddBlock(randomBlockWithValidator(t, 3, types.Hash{})))
}

func getPrevBlockHash(t *testing.T, bc *BlockChain, height uint32) types.Hash {
	prevHeader, err := bc.GetHeader(height - 1)
	assert.Nil(t, err)
	return BlockHasher{}.Hash(prevHeader)
}
