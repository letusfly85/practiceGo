/**
 * https://gobyexample.com/sha1-hashes
 *
 */

package main

import "crypto/sha1"
import "fmt"

func main() {

	s := "sha1 this string"

	h := sha1.New()

	h.Write([]byte(s))

	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
