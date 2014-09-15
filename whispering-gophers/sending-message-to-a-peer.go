/**
 * refs: http://whispering-gophers.appspot.com/talk.slide#16
 * refs: https://gist.github.com/iwanbk/2295233
 *
 * usage: go run sending-message-to-a-peer.go hoge
 *
 *
 *
 */

package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
)
import "os"

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
	strEcho := os.Args[1]
	addr, err := ioutil.ReadFile("/tmp/tcp-serv-addr")
	check(err)
	servAddr := string(addr)

	conn, err := net.Dial("tcp", servAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	var site = Site{Addr: "", Title: strEcho, URL: servAddr}
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	err = enc.Encode(site)
	if err != nil {
		log.Fatal(err)
	}

	str := b.String()
	_, err = conn.Write([]byte(str))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	reply := make([]byte, 1024)
	_, err = conn.Read(reply)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	println("reply from server=", string(reply))
}
