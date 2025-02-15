package depth_first_search

import (
	"fmt"
	"testing"
)

func TestBuildMap(t *testing.T) {
	tmpMap := BuildMap(50, 33, 3)
	fmt.Println(tmpMap)
}

func TestSaveHaDfs(t *testing.T) {
	tmpMap := BuildMap(50, 50, 15)
	targetX, targetY := SetTarget(tmpMap)
	fmt.Println("targetX", targetX, "targetY", targetY)
	for row := range tmpMap {
		for col := range tmpMap[row] {
			switch tmpMap[row][col] {
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
	book := make([][]int, len(tmpMap))
	for i := range book {
		book[i] = make([]int, len(tmpMap[i]))
	}
	book[0][0] = 1
	minStep := -1
	SaveHaDfs(tmpMap, book, 0, 0, &minStep, 0)
	fmt.Println(minStep)
	fmt.Println(book)
}
