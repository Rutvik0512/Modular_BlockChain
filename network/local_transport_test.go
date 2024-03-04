package network

import (
	"encoding/gob"
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

	msg := NewMessage(MessageTypeTx, []byte("Hello World"))
	assert.Nil(t, tra.SendMessage(trb.Addr(), msg.Bytes()))

	rpc := <-trb.(*LocalTransport).consumeCh
	message := new(Message)
	assert.Nil(t, gob.NewDecoder(rpc.Payload).Decode(message))
	assert.Equal(t, message.Data, msg.Data)
	assert.Equal(t, string(tra.Addr()), rpc.From)
}
