/**
 * refs: http://whispering-gophers.appspot.com/talk.slide#36
 *
 *
 *
 */

package main

import (
	"sync"
	"time"
)

var (
	count int
	mu    sync.Mutex // protects count
)

func main() {
	for i := 0; i < 10; i++ {
		go func() {
			for {
				mu.Lock()
				count++
				mu.Unlock()
				time.Sleep(5 * time.Millisecond)
			}
		}()
	}
	time.Sleep(time.Second)
	mu.Lock()
	println(count)
	mu.Unlock()
}
