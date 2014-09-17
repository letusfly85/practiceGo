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

	for {
		conn, err := l.Accept()
		check(err)
		go handleRequest(conn, l.Addr().String())
	}
}

func handleRequest(conn net.Conn, addr string) {
	defer conn.Close()

	buf := make([]byte, 1024)
	reqLen, err := conn.Read(buf)
	check(err)
	s := string(buf[:reqLen])

	site := new(Site)
	err = json.Unmarshal([]byte(s), &site)
	check(err)

	site.Addr = addr
	fmt.Fprintln(conn, (site.Addr + "\t" + site.Message))
	println(site.Addr, site.Message)
	io.Copy(conn, conn)
}
