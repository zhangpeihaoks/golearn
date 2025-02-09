package stack

import (
	"fmt"
	"log/slog"
	"testing"
)

func main() {
	input := "ahhha"
	inputBytes := []byte(input)

	length := len(inputBytes)

	mid := length/2 - 1

	s := NewStack[byte](mid)

	for i := 0; i <= mid; i++ {
		s.Push(inputBytes[i])
	}
	var next int
	if length%2 == 0 {
		next = mid + 1
	} else {
		next = mid + 2
	}

	fmt.Printf("%s\n", s.Data)

	for i := next; i <= length-1; i++ {
		if s.Data[s.Top] != inputBytes[i] {
			break
		}
		s.Top--
	}
	if s.Top == 0 {
		slog.Info("match")
	} else {
		slog.Info("not match")
	}

}

func Test(t *testing.T) {
	main()
}
