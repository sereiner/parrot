package server

import (
	"time"

	"github.com/sereiner/parrot/servers/rpc/codec"
	"github.com/sereiner/parrot/servers/rpc/protocol"
	"github.com/sereiner/parrot/servers/rpc/registry"
	"github.com/sereiner/parrot/servers/rpc/transport"
)

type Option struct {
	AppKey         string
	Registry       registry.Registry
	RegisterOption registry.RegisterOption
	Wrappers       []Wrapper
	ShutDownWait   time.Duration
	ShutDownHooks  []ShutDownHook

	ProtocolType  protocol.ProtocolType
	SerializeType codec.SerializeType
	CompressType  protocol.CompressType
	TransportType transport.TransportType
}

var DefaultOption = Option{
	ShutDownWait:  time.Second * 12,
	ProtocolType:  protocol.Default,
	SerializeType: codec.MessagePack,
	CompressType:  protocol.CompressTypeNone,
	TransportType: transport.TCPTransport,
}

type ShutDownHook func(s *SGServer)
