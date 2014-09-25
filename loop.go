package main

import (
	"log"
	"os/exec"
	"strconv"
)

func main() {
	url := "10.23.3.2:3000/users/1/change_name?name=hoge"
	for i := 0; i <= 10000; i++ {
		_url := url + strconv.Itoa(i)
		cmd := exec.Command("curl", "-XGET", _url)
		out, err := cmd.Output()

		println(_url)
		if err != nil {
			log.Fatal(err)
		} else {
			println(out)
		}
	}
}
