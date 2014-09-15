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
	"log"
	"net"
)

type Site struct {
	Title string
	URL   string
}

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

func main() {
	//l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	l, err := net.Listen(CONN_TYPE, ":"+CONN_PORT)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(l.Addr())
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleRequest(conn, l.Addr().String())
	}
}

func handleRequest(conn net.Conn, addr string) {
	defer conn.Close()

	buf := make([]byte, 1024)
	reqLen, err := conn.Read(buf)
	if err != nil {
		println("Error reading:", err.Error())
	}
	s := string(buf[:reqLen])

	site := new(Site)
	err = json.Unmarshal([]byte(s), &site)
	fmt.Fprintln(conn, site.Title)
	println(site.Title, addr)
	io.Copy(conn, conn)
}
