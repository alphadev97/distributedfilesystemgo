package p2p

// handshakeFunc....
type HandshakeFunc func(Peer) error

func NOPHandshakeFunc(Peer) error { return nil }
