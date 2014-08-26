/**
 * https://gobyexample.com/select
 *
 */

package main

import "time"
import "fmt"

func foo() {
	for {
		select {
		case t1 := <-time.After(time.Second):
			fmt.Println("hello", t1)
			if t1.Second()%4 == 0 {
				return
			}
		}
	}
}

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}

	foo()
	fmt.Println("end")
}
