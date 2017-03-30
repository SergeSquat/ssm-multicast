package main

import (
	"fmt"
	"net"

	"encoding/hex"
	"golang.org/x/net/ipv4"
	"log"
)

const (
	network         = "ppp0"
	mcAddr          = "239.192.111.1:9011"
	src             = "10.50.129.200"
	group           = "239.192.111.1"
	maxDatagramSize = 8192
)

func main() {
	c, err := net.ListenPacket("udp", mcAddr)

	if err != nil {
		log.Fatal(err)
	}

	defer c.Close()

	p := ipv4.NewPacketConn(c)

	ifi, err := net.InterfaceByName(network)

	if err != nil {
		log.Fatal(err)
	}

	g := net.UDPAddr{IP: net.ParseIP(group)}
	s := net.UDPAddr{IP: net.ParseIP(src)}

	err = p.JoinSourceSpecificGroup(ifi, &g, &s)

	if err != nil {
		log.Fatal(err)
	}

	for {
		b := make([]byte, maxDatagramSize)

		n, _, src, err := p.ReadFrom(b)

		if err != nil {
			log.Fatal("ReadFromUDP failed:", err)
		}

		fmt.Println(n, "bytes read from", src)
		fmt.Println(hex.Dump(b[:n]))
	}

}
