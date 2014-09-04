package main

import (
	"fmt"
	"time"
)

func timeout(t chan bool) {
	time.Sleep(5000 * 1000 * 1000)
	t <- true
}

func readString(s chan string) {
	time.Sleep(10000 * 1000 * 1000)
	msg := "This is another timeout!"
	s <- msg
}

func countDown(count chan int) {
	for i := 10; i >= 0; i-- {
		count <- i
		time.Sleep(1000 * 1000 * 1000)
	}
}

func main() {
	t := make(chan bool)
	s := make(chan string)
	c := make(chan int)

	go countDown(c)
	go readString(s)
	go timeout(t)

	for {
		select {
		case i := <-c:
			fmt.Printf("%d seconds remaining\n", i)

		case msg := <-s:
			fmt.Printf("Received: %s\n", msg)
			return

		case <-t:
			fmt.Println("Timed out\n")
			return
		}
	}
}
