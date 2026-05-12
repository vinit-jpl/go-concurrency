package main

import "fmt"

const (
	BATCH_SIZE = 5
)

func sumWorker(s []int, ch chan int) {

	sum := 0

	for _, v := range s {
		sum += v
	}

	ch <- sum

}
func checkSum(s []int) {
	sum := 0
	for _, v := range s {
		sum += v
	}

	fmt.Println("Normal sum:", sum)
}

func concurrentSum(s []int, ch chan int) {

	numberOfBatches := (len(s) + BATCH_SIZE - 1) / BATCH_SIZE
	for i := 0; i < len(s); i += BATCH_SIZE {
		end := i + BATCH_SIZE

		// prevent out of bound error
		if end > len(s) {
			end = len(s)
		}

		go sumWorker(s[i:end], ch)
	}

	res := 0

	for i := 0; i < numberOfBatches; i++ {
		res += <-ch
	}

	fmt.Println("Concurrent sum: ", res)

}
func main() {

	s := []int{438, 6849, 4501, 7028, 5351, 6414, 2240, 9939, 3352, 2425, 2418, 2966, 2714, 6533, 164, 7407, 106, 7016, 3810, 9490, 2448, 7072, 9441, 388, 1063, 6914, 5961}
	ch := make(chan int)

	concurrentSum(s, ch)

	checkSum(s)

}
