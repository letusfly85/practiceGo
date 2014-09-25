package main

import (
	"log"
	"os/exec"
	"strconv"
	"time"
)

func main() {
	url := "10.23.3.2:3000/users/1/change_name?name=hoge"
	for i := 0; i <= 10000; i++ {
		_url := url + strconv.Itoa(i)
		cmd := exec.Command("curl", "-XGET", _url)
		time.Sleep(time.Millisecond * 500)
		out, err := cmd.Output()

		log.Printf("%v", _url)
		if err != nil {
			log.Fatal(err)
		} else {
			println(out)
		}
	}
}
