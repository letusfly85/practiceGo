package main

import (
	"log"
	"net"
	"strconv"
	"sync"
	"time"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

const (
	MAX_BUFFER = 10
)

func main() {
	var _pos = 10
	//var ch = make(chan int)
	var wg sync.WaitGroup

	for i := 1; i <= MAX_BUFFER; i++ {
		pos := _pos * i
		wg.Add(1)
		go func() {
			//exec(ch, pos*i)
			exec(pos)
			wg.Done()
		}()
	}
	wg.Wait()
}

//func exec(ch chan<- int, pos int) {
func exec(pos int) {
	var rfudp int
	var err error

	/*
		defer func() {
			ch <- pos
		}()
	*/

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

	log.Printf("Receive[%d]:\t%v\n", pos, string(buf[:rfudp]))
}
