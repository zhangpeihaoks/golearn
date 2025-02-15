package depth_first_search

import (
	"fmt"
	"testing"
)

func TestNumPlusDfs(t *testing.T) {
	numList := make([]int, 10)
	book := make([]int, 10)
	total := 0
	total = NumPlusDfs(1, numList, book, total)
	fmt.Println()
	fmt.Println("total:", total/2)
}
