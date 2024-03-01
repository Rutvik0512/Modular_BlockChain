package core

import (
	"main/crypto"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignTransaction(t *testing.T) {
	tx := &Transaction{
		Data: []byte("Hello World"),
	}
	privKey := crypto.GeneratePrivateKey()

	assert.Nil(t, tx.Sign(privKey))
}

func TestVerifyTrnsaction(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	tx := &Transaction{
		Data: []byte("Hello World"),
	}

	assert.Nil(t, tx.Sign(privKey))
	assert.Nil(t, tx.Verify())

	otherPrivateKey := crypto.GeneratePrivateKey()
	tx.From = otherPrivateKey.PublicKey()

	assert.NotNil(t, tx.Verify())
}
