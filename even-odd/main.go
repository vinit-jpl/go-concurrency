package main

import (
	"fmt"
	"sync"
)

// func CalculateEvenOdd(n int, wg *sync.WaitGroup) {

// 	defer wg.Done()

// 	if n%2 == 1 {
// 		fmt.Printf("%d is odd\n", n)
// 	} else {

// 		fmt.Printf("%d is even\n", n)
// 	}
// }
// func main() {
// 	n := 20

// 	var wg sync.WaitGroup
// 	for i := 1; i <= n; i++ {
// 		wg.Add(1)
// 		go CalculateEvenOdd(i, &wg)
// 	}
// 	wg.Wait()
// }

// =========== Using Signaling =============
// func even(n int, wg *sync.WaitGroup, evenCh, oddCh chan bool) {
// 	defer wg.Done()
// 	for i := 2; i <= n; i += 2 {
// 		<-evenCh
// 		fmt.Println("even:", i)
// 		if i != n {
// 			oddCh <- true
// 		}
// 	}
// }

// func odd(n int, wg *sync.WaitGroup, evenCh, oddCh chan bool) {
// 	defer wg.Done()
// 	for i := 1; i <= n; i += 2 {
// 		<-oddCh
// 		fmt.Println("odd:", i)
// 		if i != n {
// 			evenCh <- true
// 		}
// 	}
// }

// func main() {
// 	n := 10
// 	var wg sync.WaitGroup
// 	oddCh := make(chan bool)
// 	evenCh := make(chan bool)

// 	wg.Add(2)
// 	go even(n, &wg, evenCh, oddCh)
// 	go odd(n, &wg, evenCh, oddCh)

// 	oddCh <- true

// 	wg.Wait()
// }

// ============== using struct as signal ====================

func odd(n int, oddCh, evenCh chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= n; i += 2 {

		_, ok := <-oddCh
		if !ok {
			return
		}

		fmt.Println("odd:", i)

		if i+1 <= n {
			evenCh <- struct{}{}
		} else {
			close(evenCh)
		}
	}
}

func even(n int, oddCh, evenCh chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 2; i <= n; i += 2 {

		_, ok := <-evenCh
		if !ok {
			return
		}

		fmt.Println("even:", i)

		if i+1 <= n {
			oddCh <- struct{}{}
		} else {
			close(oddCh)
		}
	}
}

func main() {
	var wg sync.WaitGroup

	n := 10

	oddCh := make(chan struct{})
	evenCh := make(chan struct{})

	wg.Add(2)

	go odd(n, oddCh, evenCh, &wg)
	go even(n, oddCh, evenCh, &wg)

	oddCh <- struct{}{}

	wg.Wait()
}
