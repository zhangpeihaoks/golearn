package sort

import (
	"context"
	"runtime"
	"slices"
	"sort"
	"sync"
)

func ParallelSort(ctx context.Context, arr []int, asc bool) {
	if len(arr) < 2 {
		return
	}
	ch := make(chan struct{})
	defer close(ch)

	go func() {

		numThreads := runtime.NumCPU()
		// 分割数组
		chunkSize := (len(arr) + numThreads - 1) / numThreads // 每个线程负责的块大小
		subArrays := make([][]int, 0, numThreads)

		for i := 0; i < len(arr); i += chunkSize {
			end := i + chunkSize
			if end > len(arr) {
				end = len(arr)
			}
			subArrays = append(subArrays, arr[i:end])
		}

		// 并行排序每个子数组
		var wg sync.WaitGroup
		for i := 0; i < len(subArrays); i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()

				sort.Ints(subArrays[i]) // 使用标准库对子数组排序

			}(i)
		}
		wg.Wait() // 等待所有 Goroutines 完成
		arr = mergeSortedArrays(subArrays)
		if !asc {
			slices.Reverse(arr)
		}
		ch <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		return
	case <-ch:
		return
	}

}

// 合并多个有序数组
func mergeSortedArrays(arrays [][]int) []int {
	if len(arrays) == 1 {
		return arrays[0]
	}

	for len(arrays) > 1 {
		merged := make([][]int, 0, (len(arrays)+1)/2)
		for i := 0; i < len(arrays); i += 2 {
			if i+1 < len(arrays) {
				merged = append(merged, mergeTwoArrays(arrays[i], arrays[i+1]))
			} else {
				merged = append(merged, arrays[i])
			}
		}
		arrays = merged
	}
	return arrays[0]
}

// 合并两个有序数组
func mergeTwoArrays(arr1, arr2 []int) []int {
	result := make([]int, len(arr1)+len(arr2))
	i, j, k := 0, 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] <= arr2[j] {
			result[k] = arr1[i]
			i++
		} else {
			result[k] = arr2[j]
			j++
		}
		k++
	}
	for i < len(arr1) {
		result[k] = arr1[i]
		i++
		k++
	}
	for j < len(arr2) {
		result[k] = arr2[j]
		j++
		k++
	}
	return result
}
