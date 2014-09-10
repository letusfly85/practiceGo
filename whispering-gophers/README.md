* Part1: reading and writing
	* site

	  http://whispering-gophers.appspot.com/talk.slide#15

	  Key,Valueを持つ構造体を用意し、標準出力の内容を受け取ったら
	  Json形式に変換して表示する簡単なプログラムです。



```Go
package main

import (
	"bufio"
	"fmt"
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
			fmt.Println("done")
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
```
