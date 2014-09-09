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
		text := scanner.Text()

		switch {
		case text == "exit":
			return

		default:

			var output = MyOutput{"Body", text}
			enc := json.NewEncoder(os.Stdout)

			err := enc.Encode(output)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
