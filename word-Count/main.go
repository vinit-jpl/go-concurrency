package main

import (
	"fmt"
	"strings"
	"sync"
)

func CountWords(text string, wg *sync.WaitGroup, resultChan chan int) {

	defer wg.Done()

	words := strings.Fields(text)

	resultChan <- len(words)
}

func main() {
	// Input strings
	texts := []string{
		"Go routines are lightweight",
		"Concurrency is powerful in Golang",
		"Channels help goroutines communicate",
		"Word counting using concurrency",
	}

	var wg sync.WaitGroup
	resultChan := make(chan int, len(texts))

	// spwan go routines

	for _, text := range texts {
		wg.Add(1)
		go CountWords(text, &wg, resultChan)
	}

	go func() {

		wg.Wait()
		close(resultChan)
	}()

	totalWords := 0

	for v := range resultChan {

		totalWords += v
	}

	fmt.Println("Total Words:", totalWords)

}
