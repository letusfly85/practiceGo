package main

import (
	"log"
	"net"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	addr, err := net.ResolveUDPAddr("udp", "localhost:8888")
	check(err)

	conn, err := net.ListenUDP("udp", addr)
	check(err)
	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		rfudp, remote, err := conn.ReadFromUDP(buf)
		check(err)

		s := string(buf[:rfudp])
		log.Printf("[%v]\t%v\n", remote, s)

		rfudp, err = conn.WriteToUDP([]byte(s), remote)
		check(err)
	}
}
