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
	"fmt"
	"log"
	"net"
)
import "os"

type Site struct {
	Title string
	URL   string
}

// go run sending-messages-to-a-peer.go "message" "localhost:3333"
func main() {

	strEcho := os.Args[1]
	servAddr := os.Args[2]
	conn, err := net.Dial("tcp", servAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	var site = Site{Title: strEcho, URL: servAddr}
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	err = enc.Encode(site)
	if err != nil {
		fmt.Println(err)
	}

	str := b.String()
	println(str)
	_, err = conn.Write([]byte(str))
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}

	reply := make([]byte, 1024)
	_, err = conn.Read(reply)
	if err != nil {
		println("Reply From server failed:", err.Error())
		os.Exit(1)
	}
	println("reply from server=", string(reply))
}
