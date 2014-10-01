/**
 *
 *
 */

package main

import (
	"bytes"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func main() {
	max_count := flag.Int("num", 10, "you can give a loop number to call curl command")
	forkPattern := flag.Bool("fork", false, "if you want to fork, give true")
	flag.Parse()

	if *forkPattern {
		daemon()
	}

	url := "10.23.3.2/twem/users/1/change_name?name=hoge"
	for i := 0; i <= *max_count; i++ {
		_url := url + strconv.Itoa(i)
		_url = _url + "&age=" + strconv.Itoa(i)
		cmd := exec.Command("curl", "-XGET", _url, "-b", "cookie.txt", "-c", "cookie.txt")
		time.Sleep(time.Millisecond * 250)
		out, err := cmd.Output()

		log.Printf("%v", _url)
		if err != nil {
			log.Fatal(err)
		} else {
			println(out)
		}
	}
}

// refs: https://github.com/astaxie/build-web-application-with-golang/blob/master/ja/ebook/12.3.md
func daemon() {
	//if *d {
	cmd := exec.Command(os.Args[0],
		"-close-fds",
		//"-addr", *addr,
		//"-call", *call,
	)
	serr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatalln(err)
	}
	err = cmd.Start()
	if err != nil {
		log.Fatalln(err)
	}
	s, err := ioutil.ReadAll(serr)
	s = bytes.TrimSpace(s)
	if bytes.HasPrefix(s, []byte("addr: ")) {
		println(string(s))
		cmd.Process.Release()
	} else {
		log.Printf("unexpected response from MarGo: `%s` error: `%v`\n", s, err)
		cmd.Process.Kill()
	}
	//}
}
