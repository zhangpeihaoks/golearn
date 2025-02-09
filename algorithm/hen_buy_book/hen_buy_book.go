package main

import (
	"fmt"
	"golearn/algorithm/quick_sort"
	"math/rand"
	"time"
)

func main() {
	buyers := 10000
	booksIsbn := make([]int, buyers)
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < buyers; i++ {
		booksIsbn[i] = rand.Intn(10000)
	}
	startTime := time.Now()

	booksIsbn = quick_sort.QuickSort(booksIsbn, true)

	for i := 0; i < len(booksIsbn)-1; i++ {
		if booksIsbn[i] != booksIsbn[i+1] {
			fmt.Printf("%v\t", booksIsbn[i])
		}
	}
	fmt.Printf("%v\n", booksIsbn[len(booksIsbn)-1])
	endTime := time.Now()

	fmt.Printf("%v\n", endTime.Sub(startTime).String())
}
