/**
 *
 *
 */

package main

import (
	"bytes"
	"code.google.com/p/gcfg"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"
)

type Config struct {
	Curl struct {
		URL string
	}
}

func check(err error) {
	if err != nil {
		log.Fatalf("error :%v\n", err)
	}
}

func main() {
	// command line 引数を受け取ります。
	max_count := flag.Int("num", 10, "curlを実行する回数を設定できます。デフォルトでは10回になります。")
	forkPattern := flag.Bool("fork", false, "バックグラウンドプロセス実行したいときにtrueを指定してください。デフォルトではfalseになります。")
	flag.Parse()

	if *forkPattern {
		daemon()
	}

	// 設定ファイルを読み込みます。
	var cfg Config
	err := gcfg.ReadFileInto(&cfg, "my-curl.conf")
	check(err)

	url := cfg.Curl.URL
	for i := 0; i <= *max_count; i++ {
		_url := url + strconv.Itoa(i)
		_url = _url + "&age=" + strconv.Itoa(i)
		cmd := exec.Command("curl", "-XGET", _url, "-b", "cookie.txt", "-c", "cookie.txt")
		time.Sleep(time.Millisecond * 250)
		out, err := cmd.Output()

		log.Printf("%v", _url)
		check(err)
		println(out)
	}
}

// refs: https://github.com/astaxie/build-web-application-with-golang/blob/master/ja/ebook/12.3.md
func daemon() {
	//if *d {
	cmd := exec.Command(os.Args[0])
	//"-close-fds",
	//"-addr", *addr,
	//"-call", *call,

	serr, err := cmd.StderrPipe()
	check(err)

	err = cmd.Start()
	check(err)

	s, err := ioutil.ReadAll(serr)
	check(err)

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
