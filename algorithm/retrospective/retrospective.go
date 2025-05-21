package retrospective

import (
	"fmt"
)

type QueenNode struct {
	Row, Col int
}

func NQueen(arr [][]int, n int) {
	queens := make([]*QueenNode, 0)
	results := make([][]*QueenNode, 0)
	backtrack(&results, queens, arr, n, 0, 0)
	if len(results) == 0 {
		fmt.Println("没有解法")
		return
	}
	fmt.Println("解法数量:", len(results))
	for index, solution := range results {
		fmt.Printf("解法%d:\t", index+1)
		for _, queen := range solution {
			fmt.Printf("(%d,%d)\t", queen.Row, queen.Col)
		}
		fmt.Println()
	}
}

func backtrack(results *[][]*QueenNode, queens []*QueenNode, arr [][]int, n, startRow, startCol int) {
	if len(queens) == n {
		solution := make([]*QueenNode, len(queens))
		copy(solution, queens)
		*results = append(*results, solution)
		return
	}

	for row := startRow; row < len(arr); row++ {

		for col := func() int {
			if row == startRow {
				return startCol
			}
			return 0
		}(); col < len(arr[row]); col++ {
			if checkInvalid(queens, row, col) {
				queens = append(queens, &QueenNode{Row: row, Col: col}) // 放置皇后
				backtrack(results, queens, arr, n, row+1, 0)            // 从当前行和下一列继续递归
				queens = queens[:len(queens)-1]                         // 回溯：移除皇后
			}
		}
	}
}

func checkInvalid(queenList []*QueenNode, row, col int) bool {
	for _, item := range queenList {
		if item.Row == row || item.Col == col || (abs(row-item.Row) == abs(col-item.Col)) {
			return false
		}
	}
	return true
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
