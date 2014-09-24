/**
 * https://gobyexample.com/goroutines
 *
 */

package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	i := 1
	var ary [100]int
	for i < 100 {
		ary[i] = 100 - i
		i += 1
	}
	fmt.Println(ary)

	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
