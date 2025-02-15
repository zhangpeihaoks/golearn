package depth_first_search

import "testing"

func TestDepthFirstSearch(t *testing.T) {
	boxCount := 3
	boxData := make([]int, boxCount)
	book := make([]int, boxCount)
	for i := 1; i <= boxCount; i++ {
		boxData = append(boxData, 0)
		book = append(book, 0)
	}
	dfs(1, boxCount, boxData, book)
}
