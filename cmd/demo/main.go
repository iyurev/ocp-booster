package main

import (
	"fmt"
	"github.com/insomniacslk/dhcp/dhcpv4"
	"github.com/insomniacslk/dhcp/dhcpv4/server4"
	"log"
	"net"
)

var (
	serverAddr = net.ParseIP("192.168.43.1")
)

func Mask24() net.IPMask {
	return net.IPv4Mask(255, 255, 255, 0)
}

func handler(conn net.PacketConn, peer net.Addr, m *dhcpv4.DHCPv4) {
	//log.Println(m.Summary())
	if m.ClientHWAddr.String() == "52:54:00:3b:da:fe" {
		if m.MessageType() == dhcpv4.MessageTypeDiscover {
			log.Println("We received a DISCOVERING message")
			log.Println(fmt.Sprintf("MAC: %s SERVER ADDR: %s YOUR ADDR: %s\n", m.ClientHWAddr.String(), m.ServerIPAddr.String(), m.YourIPAddr.String()))
			repl, err := dhcpv4.New(dhcpv4.WithMessageType(dhcpv4.MessageTypeOffer),
				dhcpv4.WithYourIP(net.ParseIP("192.168.43.100")),
				dhcpv4.WithGatewayIP(net.ParseIP("192.168.43.1")),
				dhcpv4.WithBroadcast(true),
				dhcpv4.WithNetmask(Mask24()),
				dhcpv4.WithBroadcast(true),
				dhcpv4.WithReply(m))
			if err != nil {
				log.Println(err)
			}
			_, err = conn.WriteTo(repl.ToBytes(), peer)
			if err != nil {
				log.Println(err)
			}
		}
		if m.MessageType() == dhcpv4.MessageTypeRequest {
			log.Println("We received a REQUEST message from a client")
			log.Println(fmt.Sprintf("MAC: %s SERVER ADDR: %s YOUR ADDR: %s\n", m.ClientHWAddr.String(), m.ServerIPAddr.String(), m.YourIPAddr.String()))
			log.Println(m.Summary())
			repl, err := dhcpv4.New(dhcpv4.WithMessageType(dhcpv4.MessageTypeAck),
				//dhcpv4.WithReply(m),
				dhcpv4.WithYourIP(net.ParseIP("192.168.43.100")),
				dhcpv4.WithGatewayIP(net.ParseIP("192.168.43.1")),
				dhcpv4.WithNetmask(Mask24()),
				dhcpv4.WithBroadcast(true),
				dhcpv4.WithLeaseTime(3000),
				dhcpv4.WithReply(m))
			if err != nil {
				log.Println(err)
			}
			_, err = conn.WriteTo(repl.ToBytes(), peer)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func main() {
	laddr := net.UDPAddr{
		IP:   net.ParseIP("0.0.0.0"),
		Port: dhcpv4.ServerPort,
	}

	log.Println("Start dhcp server")
	//con, err := server4.NewIPv4UDPConn("ocplab0", laddr)
	//if err != nil {
	//	log.Fatal(err)
	//}
	server, err := server4.NewServer("ocplab0", &laddr, handler, server4.WithDebugLogger())
	if err != nil {
		log.Fatal(err)
	}
	err = server.Serve()
	if err != nil {
		log.Fatal(err)
	}

}
