package depth_first_search

import (
	"fmt"
	"slices"
	"sync"
	"testing"

	"github.com/gogf/gf/v2/os/gtime"
)

func TestHorseDfs(t *testing.T) {
	n := 100
	checkBoard := make([][]int, n)
	for i := 0; i < len(checkBoard); i++ {
		checkBoard[i] = make([]int, n)
	}
	startX, startY := 0, 0
	size := len(checkBoard) * len(checkBoard[0])
	startTime := gtime.Now()
	wg := sync.WaitGroup{}

	arr1 := slices.Clone(checkBoard)
	arr2 := slices.Clone(checkBoard)

	wg.Add(2)
	go func() {
		defer wg.Done()
		HorseDfs(arr1, startX, startY, 1, size)
		fmt.Println("HorseDfs finished,process time:", gtime.Now().Sub(startTime).String())
	}()
	go func() {
		defer wg.Done()
		HorseWarnsdorff(arr2, startX, startY, 1, size)
		fmt.Println("HorseWarnsdorff finished,process time:", gtime.Now().Sub(startTime).String())
	}()

	wg.Wait()
	//for _, rows := range checkBoard {
	//	for _, col := range rows {
	//		fmt.Printf("%d \t", col)
	//	}
	//	fmt.Println()
	//}
}
