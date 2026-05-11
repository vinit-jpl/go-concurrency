package main

import (
	"fmt"
	"time"
)

func worker(id int, ch chan bool) {

	fmt.Printf("worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("worker %d done \n", id)
	ch <- true
}
func main() {

	chan1 := make(chan bool)

	for i := 0; i < 5; i++ {
		// go func(chan bool) {
		// 	worker(i, chan1)
		// }(chan1)
		// <-chan1

		go worker(i, chan1)
		<-chan1
	}
}
