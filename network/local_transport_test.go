package network

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	tra.Connect(trb)
	trb.Connect(tra)
	assert.Equal(t, tra.(*LocalTransport).peers[trb.Addr()], trb)
	assert.Equal(t, trb.(*LocalTransport).peers[tra.Addr()], tra)
}

func TestSendMessage(t *testing.T) {
	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	tra.Connect(trb)
	trb.Connect(tra)

	msg := []byte("Hello World")
	assert.Nil(t, tra.SendMessage(trb.Addr(), msg))

	rpc := <-trb.(*LocalTransport).consumeCh
	assert.Equal(t, rpc.Payload, msg)
	assert.Equal(t, string(tra.Addr()), rpc.From)
}