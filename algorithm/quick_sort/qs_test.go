package quick_sort

import (
	"log/slog"
	"math/rand"
	"runtime"
	"testing"
	"time"
)

func TestQuickSort(t *testing.T) {
	// 记录初始内存状态
	var memStatsBefore, memStatsAfter runtime.MemStats
	runtime.GC() // 手动触发 GC 避免干扰
	runtime.ReadMemStats(&memStatsBefore)

	arr := make([]int, 0)
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(10))
	}
	//slog.Info("Input arr", "arr", arr)

	startTime := time.Now()
	//arr = QuickSort(arr, false)
	QuickSort2(arr, 0, len(arr)-1)
	endTime := time.Now()

	// 记录排序后内存状态
	runtime.ReadMemStats(&memStatsAfter)

	//slog.Info("Output arr", "arr", arr)

	// 打印结果
	slog.Info("Memory usage",
		"allocated(MB)", (memStatsAfter.TotalAlloc-memStatsBefore.TotalAlloc)/1024/1024,
		"heap_inuse(MB)", (memStatsAfter.HeapInuse-memStatsBefore.HeapInuse)/1024/1024,
	)
	slog.Info("Duration", "time", endTime.Sub(startTime).String())
}
