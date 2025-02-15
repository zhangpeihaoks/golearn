package depth_first_search

import (
	"fmt"
	"math/rand"
	"time"
)

func BuildMap(row, col int, obstacleCount int) [][]int {
	tmpMap := make([][]int, row)
	for i := range tmpMap {
		tmpMap[i] = make([]int, col)
	}
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < obstacleCount; i++ {
		for {
			obX := rand.Intn(len(tmpMap))
			obY := rand.Intn(len(tmpMap[obX]))
			if tmpMap[obX][obY] == 0 {
				tmpMap[obX][obY] = 1
				break
			}
		}
	}
	return tmpMap
}

func SetTarget(theMap [][]int) (int, int) {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		targetX := rand.Intn(len(theMap))
		targetY := rand.Intn(len(theMap[targetX]))
		if theMap[targetX][targetY] == 0 {
			theMap[targetX][targetY] = 2
			return targetX, targetY
		}
	}
}

var next = [4][2]int{
	{0, 1}, {0, -1}, {1, 0}, {-1, 0},
}

func SaveHaDfs(theMap, book [][]int, x, y int, minStep *int, step int) {
	if theMap[x][y] == 2 {
		if step < *minStep || *minStep == -1 {
			*minStep = step
		}
		list := make([]string, 0)
		for row := range book {
			for col := range book[row] {
				if book[row][col] != 0 {
					list = append(list, fmt.Sprintf("(%d,%d)", row, col))
				}
			}
		}
		for _, data := range list {
			fmt.Printf(data)
		}
		fmt.Println()
		return
	}
	var tx, ty int
	for i := 0; i < 4; i++ {
		tx = x + next[i][0]
		ty = y + next[i][1]
		if tx < 0 || tx >= len(theMap) || ty < 0 || ty >= len(theMap[tx]) {
			continue
		}

		if theMap[tx][ty] != 1 && book[tx][ty] == 0 {
			book[tx][ty] = 1

			SaveHaDfs(theMap, book, tx, ty, minStep, step+1)
			book[tx][ty] = 0
		}
	}
}
