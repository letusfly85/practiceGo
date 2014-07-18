/**
 * https://gobyexample.com/multiple-return-values
 *
 */

package main

import "fmt"

func vals() (int, int, string) {
	return 3, 7, "abc"
}

func main() {

	a, b, c := vals()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	_, _, d := vals()
	fmt.Println(d)
}
