/**
 * https://gobyexample.com/closing-channels
 *
 */

package main

import "time"
import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)

			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for k := 1; k <= 3; k++ {
		jobs <- k
		fmt.Println("sent job", k)
		time.Sleep(time.Second * 1)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done
}
