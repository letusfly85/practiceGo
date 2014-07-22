/**
 * https://gobyexample.com/non-blocking-channel-operations
 *
 */

package main

import "time"
import "fmt"

func main() {
	messages := make(chan string, 10)
	signals := make(chan bool)

	time.Sleep(time.Second * 1)

	msg := "hi"
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	msg = "hihi"
	select {
	case messages <- msg:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
