/**
 * refs: http://whispering-gophers.appspot.com/talk.slide#16
 * refs: https://gist.github.com/iwanbk/2295233
 *
 * usage: go run sending-message-to-a-peer.go hoge
 *
 */

package myserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"

	"code.google.com/p/whispering-gophers/util"
)

func peer2peer(conn net.Conn, addr string) {
	defer conn.Close()

	for {
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
	}
}

func peer() {
	l, err := util.Listen()
	check(err)

	servId := os.Args[1]
	fileName := "/tmp/tcp-peer-" + servId

	err = ioutil.WriteFile(fileName, ([]byte(l.Addr().String())), 0644)
	check(err)
	defer l.Close()

	println("peer start")
	for {
		conn, err := l.Accept()
		check(err)
		go peer2peer(conn, l.Addr().String())
	}
}
