/**
 * refs:
 *  http://whispering-gophers.appspot.com/talk.slide#12
 *
 *
 */

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Site struct {
	Title string
	URL   string
}

var sites = []Site{
	{"The Go Programming Language", "http://golang.org"},
	{"Google", "http://google.com"},
}

func main() {
	enc := json.NewEncoder(os.Stdout)
	for _, s := range sites {
		err := enc.Encode(s)
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
	}
}
