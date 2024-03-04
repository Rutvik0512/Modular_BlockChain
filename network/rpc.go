package network

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"main/core"
)

type MessageType byte

const (
	MessageTypeTx MessageType = 0x1
	MeesageTypeBlock
)

type RPC struct {
	From    string
	Payload io.Reader
}

type Message struct {
	Header MessageType
	Data   []byte
}

func NewMessage(t MessageType, data []byte) *Message {
	return &Message{
		Header: t,
		Data:   data,
	}
}

func (msg *Message) Bytes() []byte {
	buf := &bytes.Buffer{}
	gob.NewEncoder(buf).Encode(msg)
	return buf.Bytes()
}

type RPCHandler interface {
	HandleRPC(rpc RPC) error
}

type DefaultRPCHandler struct {
	p RPCProcessor
}

func NewDefaultRPCHandler(p RPCProcessor) *DefaultRPCHandler {
	return &DefaultRPCHandler{
		p: p,
	}
}

func (h *DefaultRPCHandler) HandleRPC(rpc RPC) error {
	msg := &Message{}
	if err := gob.NewDecoder(rpc.Payload).Decode(msg); err != nil {
		return fmt.Errorf("failed to Decode message From (%s): %s", rpc.From, err)
	}

	switch msg.Header {
	case MessageTypeTx:
		tx := new(core.Transaction)
		if err := tx.Decode(core.NewGobTxDecoder(bytes.NewReader(msg.Data))); err != nil {
			return err
		}
		return h.p.ProcessTransaction(NetAddr(rpc.From), tx)

	default:
		return fmt.Errorf("invalid message header %d", msg.Header)
	}

}

type RPCProcessor interface {
	ProcessTransaction(NetAddr, *core.Transaction) error
}
