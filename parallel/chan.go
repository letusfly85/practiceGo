/**
 * channelの学習用ソースコード
 *
 * refs:
 *  http://qiita.com/tenntenn/items/686a75e11e8dcd9912ec
 *
 */
package main

import "fmt"

func main() {
	m := make(map[int]string)
	m[2] = "First Value"
	c := make(chan bool, 1)
	go func() {
		m[2] = "Second Value"
		c <- true
	}()
	_ = <-c
	fmt.Printf("%s\n", m[2])

	ch1 := make(chan int, 10)
	fmt.Println(cap(ch1))

	ch2 := make(chan int)
	fmt.Println(cap(ch2))

	ch3 := make(chan struct{}, 10)
	ch3 <- struct{}{}

	fmt.Println("capture:", cap(ch3), "length:", len(ch3))

	<-ch3
	fmt.Println("capture:", cap(ch3), "length:", len(ch3))
}
