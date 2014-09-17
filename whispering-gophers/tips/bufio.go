package main

import (
	"fmt"
	"strings"
)
import "log"
import "bufio"

const input = `A haiku is more
Than just a collection of
Well-formed syllables
`

func main() {
	s := bufio.NewScanner(strings.NewReader(input))
	for s.Scan() {
		fmt.Println(s.Text())
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}
