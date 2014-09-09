package main

import (
	"bufio"
	"os"
)
import "encoding/json"
import "log"

type MyOutput struct {
	Key   string
	Value string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var output = MyOutput{"Body", scanner.Text()}
		enc := json.NewEncoder(os.Stdout)

		err := enc.Encode(output)
		if err != nil {
			log.Fatal(err)
		}
	}
}
