package main

import (
	"log"
	"net"
	"strconv"
	"time"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
}

func exec(ch chan<- int, pos int) {
	var rfudp int
	var err error

	defer func() {
		ch <- pos
	}()

	remote, err := net.ResolveUDPAddr("udp", "localhost:8888")
	check(err)

	conn, err := net.DialUDP("udp", nil, remote)
	check(err)

	conn.SetDeadline(time.Now().Add(5 * time.Second))
	defer conn.Close()

	s := "user\t" + strconv.Itoa(pos)

	rfudp, err = conn.Write([]byte(s))
	check(err)

	buf := make([]byte, 1024)
	rfudp, err = conn.Read(buf)
	check(err)

	log.Println("Receive[%d]:\t%v\n", pos, string(buf[:rfudp]))
}
