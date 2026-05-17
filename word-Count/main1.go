package main

import (
	"fmt"
	"strings"
	"time"
)

func wordCount(text string, ch chan int, done chan bool) {

	fmt.Printf("processing string %v \n", text)
	time.Sleep(time.Second)
	word := strings.Fields(text)

	size := len(word)

	ch <- size
	fmt.Printf("word count %d : \n", size)
	done <- false
}
