package dynamic_programming

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestArrayPlus(t *testing.T) {
	arr := make([]int, 0)
	rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(100))
	}
	fmt.Println(arr)
	arrPlus(arr)
}

func TestCountSavingWays(t *testing.T) {
	ways := countSavingWays(20)
	fmt.Println(ways)
}
