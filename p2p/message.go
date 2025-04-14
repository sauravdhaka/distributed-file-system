package p2p

import "net"

// RPC holds any arbitary data thta is being sent over the
// each transport between nods in the network
type RPC struct {
	From    net.Addr
	Payload []byte
}
