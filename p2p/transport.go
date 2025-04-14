package p2p

// Peer is an interface represents the remote node (remote connection)
type Peer interface {
	Close() error
}

// Transport is anything that handels the communication
// between the nodes in the network . This can be of the
// form (TCP , UDP , websoctest, ....)
type Transport interface {
	ListenAndAccept() error
	Consume() <-chan RPC
}
