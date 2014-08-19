/**
 * http://golang.org/pkg/path/filepath/#Abs
 *
 *
 */

package main

import "fmt"
import "path/filepath"

func main() {

	path := "c:\\learning\\go\\hoge.go"

	abs, err := filepath.Abs(path)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(abs)

	base := filepath.Base(path)
	fmt.Println(base)
}
