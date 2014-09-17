/**
 * refs: http://whispering-gophers.appspot.com/talk.slide#16
 * refs: https://gist.github.com/iwanbk/2295233
 *
 * usage: go run sending-message-to-a-peer.go hoge
 *
 */

package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"sync"
)
import "os"

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

var (
	msgCh chan string
)

func readMessage(msgCh chan string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		msgCh <- scanner.Text()
	}
}

func reactiveMessage(msgCh chan string, conn net.Conn, servAddr string) {
	for {
		select {
		case message, ok := <-msgCh:
			if !ok {
				os.Exit(1)
			}
			var site = Site{Addr: "", Message: message, URL: servAddr}
			var b bytes.Buffer
			enc := json.NewEncoder(&b)
			err := enc.Encode(site)
			check(err)

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
	}
}

func main() {
	servId := os.Args[1]
	fileName := "/tmp/tcp-serv-" + servId

	addr, err := ioutil.ReadFile(fileName)
	check(err)
	servAddr := string(addr)

	conn, err := net.Dial("tcp", servAddr)
	check(err)
	defer conn.Close()
	msgCh := make(chan string, 10)

	go reactiveMessage(msgCh, conn, servAddr)
	readMessage(msgCh)
}
