package dhcp

import (
	"github.com/insomniacslk/dhcp/dhcpv4"
	"github.com/insomniacslk/dhcp/dhcpv4/server4"
	"log"
	"net"
)

type Server struct {
	dhcpsrv *server4.Server
}

var (
	listenerAddr  = "10.10.10.21"
	listenerIface = "wlp1s0"
)

func (srv *Server) Serve() {
	if err := srv.dhcpsrv.Serve(); err != nil {
		log.Fatal(err)
	}
}

func handler(conn net.PacketConn, peer net.Addr, m *dhcpv4.DHCPv4) {
	// this function will just print the received DHCPv6 message, without replying
	log.Print(m.Summary())
}

func NewServer() (*Server, error) {
	laddr := &net.UDPAddr{
		IP:   net.ParseIP(listenerAddr),
		Port: dhcpv4.ServerPort,
	}
	srv, err := server4.NewServer(listenerIface, laddr, handler)
	if err != nil {
		return nil, err
	}
	return &Server{
		dhcpsrv: srv,
	}, nil
}
