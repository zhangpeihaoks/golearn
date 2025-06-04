package test

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := make([]int, 10000000)
	for i := range arr {
		arr[i] = rand.Intn(100000)
	}
	quickSort(arr, 0, len(arr)-1)
	//fmt.Println(arr)
}

func TestGetTag(t *testing.T) {
	fmt.Println(GetTag())
}
