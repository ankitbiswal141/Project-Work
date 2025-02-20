package p2p

import (
	"fmt"
	"net"
	"sync"
)

//TCPPeer represents the remote mode over a TCP established connection
type TCPPeer struct{
	//conn is the underlying connection of the peer
	conn net.Conn

	/*
	 * If we dial and retrieve a conn => outbound == true
	 * If we listen and retrieve a conn => inbound == true
	 */
	outbound bool
}

func NewTCPPeer(conn net.Conn,outbound bool) *TCPPeer{
	return &TCPPeer{
		conn: 		conn,
		outbound: outbound,
	}
}

type TCPTransport struct {
	ListenAddress string
	Listener      net.Listener

	mu    sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		ListenAddress: listenAddr,
	}
}

func (t *TCPTransport) ListenAccept() error {

	var err error

	t.Listener, err := net.Listen("tcp", t.ListenAddress)
	if err != nil {
		return err
	}

	go t.acceptLoop()

	return nil
}

func (t *TCPTransport) acceptLoop() {
	for {
		conn, err := t.Listener.Accept()

		if err != nil {
			fmt.Printf("TCP accept error: %s\n", err)
		}

		go t.handleConn(conn)
	}
}

func(t *TCPTransport) handleConn(conn net.Conn){
	peer:= NewTCPPeer(conn,true)

	fmt.Printf("New incomming connection %+v\n",peer)
}
