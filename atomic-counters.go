/**
 * https://gobyexample.com/atomic-counters
 *
 */

package main

import "time"
import "fmt"
import "sync/atomic"
import "runtime"

func main() {

	var ops uint64 = 0

	for i := 0; i < 50; i++ {
		go func() {
			for {

				atomic.AddUint64(&ops, 1)

				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}
