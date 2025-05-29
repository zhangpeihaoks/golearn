package depth_first_search

import (
	"fmt"
	"sort"
)

var canTouchList = [][]int{
	{-1, -2}, {-2, -1}, {-2, 1}, {-1, 2},
	{1, 2}, {2, 1}, {2, -1}, {1, -2},
}

type stepInfo struct {
	X, Y   int
	Degree int
}

// 马踏棋盘
func HorseDfs(chessboard [][]int, horseX, horseY, step, size int) {
	chessboard[horseX][horseY] = step
	if step == size {
		printChessboard(chessboard, "HorseDfs")
		return
	}

	for _, canTouch := range canTouchList {
		checkX := horseX + canTouch[0]
		checkY := horseY + canTouch[1]
		if CheckTouch(chessboard, checkX, checkY) {
			HorseDfs(chessboard, checkX, checkY, step+1, size)
			//chessboard[checkX][checkY] = 0
		}
	}

}

func HorseWarnsdorff(chessboard [][]int, horseX, horseY, step, size int) {
	chessboard[horseX][horseY] = step
	if step == size {
		printChessboard(chessboard, "HorseWarnsdorff")
		return
	}
	nextSteps := GetNextSteps(chessboard, horseX, horseY)
	sort.Slice(nextSteps, func(i, j int) bool {
		return nextSteps[i].Degree < nextSteps[j].Degree
	})

	for _, nextStep := range nextSteps {
		HorseWarnsdorff(chessboard, nextStep.X, nextStep.Y, step+1, size)
		//chessboard[nextStep.X][nextStep.Y] = 0
	}
}

func GetNextSteps(chessboard [][]int, horseX, horseY int) []stepInfo {
	steps := make([]stepInfo, 0)
	for _, canTouch := range canTouchList {
		checkX := horseX + canTouch[0]
		checkY := horseY + canTouch[1]
		if CheckTouch(chessboard, checkX, checkY) {
			steps = append(steps, stepInfo{
				X:      checkX,
				Y:      checkY,
				Degree: CalculateDegree(chessboard, checkX, checkY),
			})
		}
	}
	return steps
}

func CalculateDegree(chessboard [][]int, x, y int) int {
	count := 0
	for _, canTouch := range canTouchList {
		checkX := x + canTouch[0]
		checkY := y + canTouch[1]
		if CheckTouch(chessboard, checkX, checkY) {
			count++
		}
	}
	return count
}

func CheckTouch(chessboard [][]int, checkX, checkY int) bool {
	return checkX >= 0 && checkY >= 0 &&
		checkX < len(chessboard) && checkY < len(chessboard[0]) &&
		chessboard[checkX][checkY] == 0
}

func printChessboard(chessboard [][]int, name string) {
	fmt.Println(name)
	// 输出完整路径
	for _, row := range chessboard {
		for _, col := range row {
			fmt.Printf("%2d\t", col)
		}
		fmt.Println()
	}
	fmt.Println()
}
