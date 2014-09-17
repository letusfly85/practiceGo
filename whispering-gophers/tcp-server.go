/**
 * refs: http://whispering-gophers.appspot.com/talk.slide#16
 * refs: https://coderwall.com/p/wohavg
 * refs: http://d.hatena.ne.jp/taknb2nch/20140210/1392044307
 * refs: http://ramtiga.hatenablog.jp/entry/2013/11/06/184113
 *
 *
 */

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"sync"

	"code.google.com/p/whispering-gophers/util"
)

type Peers struct {
	m  map[string]chan<- Message
	mu sync.RWMutex
}

type Message string

type Site struct {
	Addr    string
	Message string
	URL     string
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	l, err := util.Listen()
	check(err)

	servId := os.Args[1]
	fileName := "/tmp/tcp-serv-" + servId
	err = ioutil.WriteFile(fileName, ([]byte(l.Addr().String())), 0644)
	check(err)
	defer l.Close()

	fileName = "/tmp/tcp-peer-" + servId
	addr, err := ioutil.ReadFile(fileName)
	check(err)
	servAddr := string(addr)

	clConn, err := net.Dial("tcp", servAddr)
	check(err)

	//TODO: define channel and give handleRequest
	for {
		conn, err := l.Accept()
		check(err)
		go handleRequest(conn, clConn, l.Addr().String())
	}

	//TODO: use channel above and get a message from it and give new method
	// call peer server

}

func handleRequest(conn net.Conn, clConn net.Conn, addr string) {
	// server
	defer conn.Close()
	buf := make([]byte, 1024)
	reqLen, err := conn.Read(buf)
	check(err)
	s := string(buf[:reqLen])

	site := new(Site)
	err = json.Unmarshal([]byte(s), &site)
	check(err)

	site.Addr = addr
	println(site.Addr, site.Message)
	fmt.Fprintln(conn, (site.Addr + "\t" + site.Message))
	io.Copy(conn, conn)

	// client
	defer clConn.Close()
	_, err = clConn.Write([]byte(s))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	reply := make([]byte, 1024)
	_, err = clConn.Read(reply)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	io.Copy(clConn, clConn)
}
