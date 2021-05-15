package main

import (
	"github.com/iyurev/ocp-booster/pkg/dhcp"
	"log"
)

func main() {
	dhcpServer, err := dhcp.NewServer()
	if err != nil {
		log.Fatal(err)
	}
	dhcpServer.Serve()
}
