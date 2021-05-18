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
	listenerAddr  = "0.0.0.0"
	listenerIface = "ocplab0"
)

func Mask24() net.IPMask {
	return net.IPv4Mask(255, 255, 255, 0)
}

func (srv *Server) Serve() {
	log.Println("DHCP server starting to serve requests")
	if err := srv.dhcpsrv.Serve(); err != nil {
		log.Fatal(err)
	}
}

func handler(conn net.PacketConn, peer net.Addr, m *dhcpv4.DHCPv4) {
	log.Println("This's DHCP v4 handler function")
	log.Print(m.Summary())
	return
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
