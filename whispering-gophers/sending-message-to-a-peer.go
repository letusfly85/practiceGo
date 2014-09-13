/**
 * refs: http://whispering-gophers.appspot.com/talk.slide#16
 * refs: https://gist.github.com/iwanbk/2295233
 *
 * usage: go run sending-message-to-a-peer.go hoge
 *
 */

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
)
import "os"

func main() {

	strEcho := os.Args[1]

	servAddr := "localhost:3333"
	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	if err != nil {
		println("Resolve TCP address failed:", err.Error())
		os.Exit(1)
	}
	println(tcpAddr.IP)
	println(tcpAddr.Port)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	d := map[string]string{"message": strEcho}
	err = enc.Encode(d)
	if err != nil {
		fmt.Println(err)
	}
	str := b.String()

	_, err = conn.Write([]byte(str))
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}

	reply := make([]byte, 1024)
	_, err = conn.Read(reply)
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}
	println("reply from server=", string(reply))
}
