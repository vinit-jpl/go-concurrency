package main

import "fmt"

func Sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}

	c <- sum
}
func ChannelSum(s []int, batchCount int) int {

	c := make(chan int)

	numOfSumBatch := len(s) / batchCount

	for i := 0; i < numOfSumBatch; i++ {
		go Sum(s[i*batchCount:(i+1)*batchCount], c)

	}

	chanSum := 0

	for i := 0; i < numOfSumBatch; i++ {
		fmt.Printf("chan sum is %d \n", chanSum)
		chanSum += <-c
	}

	return chanSum

}

func main() {

	const BATCH_COUNT = 5

	s := []int{1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, 17711, 28657, 46368, 75025, 121393}
	fmt.Println(len(s))
	fmt.Println("Sum from channel: ", ChannelSum(s, BATCH_COUNT))

}
