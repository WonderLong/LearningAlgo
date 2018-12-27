package client

import (
	"log"
	"net/rpc"
)

type RPCClient struct {
	Addr    string
	Network string
	client  *rpc.Client
}

func (c *RPCClient) Start() {
	var e error
	c.client, e = rpc.DialHTTP(c.Network, c.Addr)
	if e != nil {
		log.Fatal("Cannot connect to addr: ", c.Addr)
	}
}

type InvokeRequest struct {
	Method string
}

func (c RPCClient) Invoke(m string, param interface{}, repl interface{}) {
	c.client.Call(m, param, repl)
}
