package main

import (
	"fmt"
)

func main2() {
	ch := make(chan struct{})
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Goroutine that prints elements from `arr` after being signaled by the main goroutine.
	go func() {
		for _, v := range arr {
			<-ch           // Wait for a signal from the main goroutine
			fmt.Println(v) // Print the current element
		}
		close(ch) // Close the channel when done to prevent deadlock
	}()

	// Main goroutine prints elements and signals the worker goroutine
	for _, v := range arr {
		fmt.Println(v)   // Print the current element
		ch <- struct{}{} // Signal the worker goroutine to proceed
	}
}
