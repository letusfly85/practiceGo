/**
 * refs: http://whispering-gophers.appspot.com/talk.slide#16
 * refs: https://coderwall.com/p/wohavg
 * refs: http://d.hatena.ne.jp/taknb2nch/20140210/1392044307
 * refs: http://ramtiga.hatenablog.jp/entry/2013/11/06/184113
 *
 *
 */

package myserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"

	"code.google.com/p/whispering-gophers/util"
)

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
	peer, err := net.Dial("tcp", servAddr)
	check(err)
	defer peer.Close()

	msgCh := make(chan Site, 100)
	for {
		go func() {
			for {
				select {
				case site, ok := <-msgCh:
					if !ok {
						println("err!")
						os.Exit(1)
					}

					var b bytes.Buffer
					enc := json.NewEncoder(&b)
					err = enc.Encode(site)
					check(err)

					str := b.String()
					_, err = peer.Write([]byte(str))
					check(err)
				}
			}
		}()
		conn, err := l.Accept()
		check(err)
		go handleRequest(conn, msgCh, l.Addr().String())
	}
}

func handleRequest(conn net.Conn, msgCh chan Site, addr string) {
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
		println(site.Addr, site.Message)
		fmt.Fprintln(conn, (site.Addr + "\t" + site.Message))

		msgCh <- *site
	}
}
