package retrospective

import (
	"testing"
)

func TestNQueen(t *testing.T) {
	n := 10
	arr := make([][]int, n)
	for i := 0; i < len(arr); i++ {
		arr[i] = make([]int, n)
	}
	NQueen(arr, 10)
}
