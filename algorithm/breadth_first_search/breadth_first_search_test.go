package breadth_first_search

import (
	"fmt"
	"testing"
)

func TestBFSSaveHa(t *testing.T) {
	theMap := BuildMap(500, 500, 0)
	targetX, targetY := SetTarget(theMap)
	fmt.Println("targetX", targetX, "targetY", targetY)
	for row := range theMap {
		for col := range theMap[row] {
			switch theMap[row][col] {
			case 0:
				fmt.Printf("0 ")
			case 1:
				fmt.Printf("1 ")
			case 2:
				fmt.Printf("2 ")
			}
		}
		fmt.Println()
	}
	book := make([][]int, len(theMap))
	for row := range book {
		book[row] = make([]int, len(theMap[row]))
	}
	step := BfsSaveHa(theMap, book, 0, 0)
	fmt.Println("step:", step)
}
