package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeypair_Sign_Verify_Sucess(t *testing.T) {
	privateKey := GeneratePrivateKey()
	pubKey := privateKey.PublicKey()

	msg := []byte("hello world")
	sig, err := privateKey.Sign(msg)
	assert.Nil(t, err)

	assert.True(t, sig.Verify(pubKey, msg))
}

func TestKeypair_Sign_Verify_Fail(t *testing.T) {
	privateKey := GeneratePrivateKey()
	pubKey := privateKey.PublicKey()

	msg := []byte("hello world")
	sig, err := privateKey.Sign(msg)

	if err != nil {
		panic(err)
	}

	otherPrivateKey := GeneratePrivateKey()
	otherPublicKey := otherPrivateKey.PublicKey()

	assert.False(t, sig.Verify(otherPublicKey, msg))
	assert.False(t, sig.Verify(pubKey, []byte("xxxxxx")))
}
