package p2p

import "net"

// Peer is an interface represents the remote node (remote connection)
type Peer interface {
	Send([]byte) error
	RemoteAddr() net.Addr
	Close() error
}

// Transport is anything that handels the communication
// between the nodes in the network . This can be of the
// form (TCP , UDP , websoctest, ....)
type Transport interface {
	Dail(string) error
	ListenAndAccept() error
	Consume() <-chan RPC
	Close() error
}
