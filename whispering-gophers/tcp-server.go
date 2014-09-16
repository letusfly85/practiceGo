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

	"code.google.com/p/whispering-gophers/util"
)

type Site struct {
	Addr  string
	Title string
	URL   string
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	l, err := util.Listen()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(l.Addr())
	err = ioutil.WriteFile("/tmp/tcp-serv-addr", ([]byte(l.Addr().String())), 0644)
	check(err)
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
	site.Addr = addr
	fmt.Fprintln(conn, (site.Addr + "\t" + site.Title))
	println(site.Addr, site.Title)
	io.Copy(conn, conn)
}
