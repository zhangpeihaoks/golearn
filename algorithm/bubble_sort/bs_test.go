package bubble_sort

import (
	"log/slog"
	"math/rand"
	"testing"
	"time"
)

func TestBubbleSort(t *testing.T) {
	arr := make([]int, 0)
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(10))
	}
	slog.Info("Input arr", "arr", arr)
	startTime := time.Now()
	arr = BubbleSort(arr, false)
	endTime := time.Now()
	slog.Info("duration =", "time:", endTime.Sub(startTime))
	slog.Info("Output arr", "arr", arr)
}
